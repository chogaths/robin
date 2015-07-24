package netfw

import (
	"core/dbfw/mysql"
	"core/pbcodec"
	"flag"
	"github.com/davecheney/profile"
	"log"
	"os"
	"path/filepath"
	"protos/coredef"
	"runtime"
	"strconv"
	"strings"
)

var (
	// 命令行传入
	SvcIndex int32 = 1

	// 可执行名称(无后缀)
	SvcName string

	// 是否使用外部配置文件
	UseConfigFile bool = true

	// 实例化的peer
	peerMap map[string]IPeer = make(map[string]IPeer)

	// 宏表
	macroMap map[string]string = make(map[string]string)

	// 退出信号
	exitChan chan bool = make(chan bool)

	consoleEventArray []func(string) = make([]func(string), 0)

	DBExec dbfw.IDBExecutor

	SvcConfig coredef.ServiceConfig

	// 正在测试性能
	profiling interface {
		Stop()
	}
)

// 等待服务器结束
func WaitForExit() {

	if len(peerMap) == 0 {
		log.Println("no peer running, exit!")
		return
	}

	<-exitChan

	exitProfile()
}

// 手动退出
func Exit() {
	exitChan <- true
}

// 获取频道定义
func FindChannelDefine(name string) *coredef.ChannelDefine {

	for _, cdef := range SvcConfig.Channel {

		if cdef.GetName() == name {
			return cdef
		}
	}

	return nil
}

// 使用svcid获取端
func GetPeerByID(svcid string) IPeer {
	if v, ok := peerMap[svcid]; ok {
		return v
	}
	return nil
}

// 查找合适的端
func findPeerByType(name string, peertype string) IPeer {
	for _, p := range peerMap {
		if p.GetDefine().GetName() == name && p.GetType() == peertype {
			return p
		}
	}

	return nil
}

// 查找合适的端
func FindPeer(name string) IPeer {
	for _, p := range peerMap {
		if p.GetDefine().GetName() == name {
			return p
		}
	}

	return nil
}

// 查找连接器
func FindConnector(name string) IPeer {
	return findPeerByType(name, "connector")
}

// 查找接收器
func FindAcceptor(name string) IPeer {
	return findPeerByType(name, "acceptor")
}

// 遍历连接器
func IteratePeer(callback func(IPeer)) {
	for _, v := range peerMap {
		callback(v)
	}
}

// 初始化DB
func initDB() {

	dsn := SvcConfig.GetDB().GetDSN()

	if dsn == "" {
		return
	}

	dbConfig := SvcConfig.GetDB()

	if !dbConfig.GetEnable() {
		return
	}

	DBExec = dbfw.NewDBExecutor()

	log.Printf("openning db %s conn: %d", dsn, dbConfig.GetConnCount())

	err := DBExec.Open(dsn, dbConfig.GetShowOperate(), int(dbConfig.GetConnCount()))

	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	log.Printf("db ready")

}

func initChannel() {

	// 添加侦听器及组件
	for _, param := range SvcConfig.GetPeer() {
		addPeer(param)
	}

	// 先初始化组件
	initPeerComponent()

	// 再初始化网络
	initPeerNetwork()
}

func initPeerComponent() {

	for _, p := range peerMap {

		param := p.GetDefine()

		for _, name := range param.GetComponent() {

			p.GetComponent(name)
		}
	}

}

func initPeerNetwork() {

	for _, p := range peerMap {

		log.Printf("%s start network: %s", p.GetID(), p.GetAddress())

		if !p.GetDefine().GetManualStart() {
			p.(IPeerStarter).Start()
		}

	}
}

func addPeer(def *coredef.PeerDefine) {

	switch def.GetType() {
	case "connector":
		if def.GetPeerCount() > 1 {

			var i int32
			for i = 1; i <= def.GetPeerCount(); i++ {
				rawAddPeer(def, i)
			}

		} else {

			rawAddPeer(def, def.GetPeerIndex())
		}

	case "acceptor":
		rawAddPeer(def, SvcIndex)
	}

}

func rawAddPeer(def *coredef.PeerDefine, peerIndex int32) {
	/// 运行时信息
	peer := NewPeerData(def, pbcodec.GetInterface())
	peer.Address = def.GetAddress()
	peer.NotifyAddress = def.GetNotifyAddress()
	peer.Index = peerIndex
	peer.ID = MakeServiceID(def.GetType(), def.GetName(), peerIndex)

	// 创建peer实现
	p := CreatePeer(def.GetImplementor(), def.GetType(), peer)

	if p == nil {
		log.Printf("peer %s implementor not found, %s", peer.ID, def.GetImplementor())
		return
	}

	// 记录
	peerMap[peer.ID] = p.(IPeer)

	log.Printf("peer %s ready", peer.ID)
}

func initLog() {

	logdef := SvcConfig.GetLog()

	if !logdef.GetEnable() {
		return
	}

	var final string
	if logdef.GetFileName() == "" {
		final = SvcName + ".log"
	} else {
		final = logdef.GetFileName()
	}

	log.Printf("logfile: %s", final)

	logfile, err := os.OpenFile(final, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
		return
	}

	log.SetOutput(logfile)
}

// 测试阶段的路径
const DefaultConfigPath string = "../cfg"

// gm tool need this
var SkipCommandLine bool

type svcParam struct {
	configStr  string
	configFile string
}

func fetchCommandLine() (param svcParam) {

	if !SkipCommandLine {

		// 读取命令行参数
		var paramSvcFile = flag.String("svcfile", "", "extra run script file")
		var paramSvcCfg = flag.String("svccfg", "", "extra run script")
		var paramServiceIndex = flag.String("svcindex", "1", "service index")

		flag.Parse()

		// 构建索引号
		svcindex, _ := strconv.Atoi(*paramServiceIndex)
		SvcIndex = (int32)(svcindex)

		param.configStr = *paramSvcCfg
		param.configFile = *paramSvcFile
	}

	// 构建服务器名
	basename := filepath.Base(os.Args[0])
	SvcName = strings.TrimSuffix(basename, filepath.Ext(basename))

	return
}

// http://saml.rilspace.org/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof#rest
func initProfile() {

	if !SvcConfig.Prof.GetCPU() || !SvcConfig.Prof.GetMem() || !SvcConfig.Prof.GetBlock() {
		return
	}

	cfg := &profile.Config{
		CPUProfile:   SvcConfig.Prof.GetCPU(),
		MemProfile:   SvcConfig.Prof.GetMem(),
		BlockProfile: SvcConfig.Prof.GetBlock(),
	}

	profiling = profile.Start(cfg)
}

func exitProfile() {
	if profiling != nil {
		profiling.Stop()
	}
}

// 初始化
func Start() {

	param := fetchCommandLine()

	initConfigEnv(SvcName, param)

	initProfile()

	initLog()

	log.Printf("start %s#%d", SvcName, SvcIndex)

	log.Printf("GOMAXPROCS: %v", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	initDB()

	initChannel()
}
