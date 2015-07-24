package cache

type IDataSaver interface {
	SaveData() int
	IsDirty() bool
}

type DataSet struct {
	valueMap map[int64]interface{}
}

func (self *DataSet) AddData(key int64, value interface{}) {
	self.valueMap[key] = value
}

func (self *DataSet) RemoveData(key int64) {
	delete(self.valueMap, key)
}

func (self *DataSet) GetData(key int64) interface{} {
	if v, ok := self.valueMap[key]; ok {
		return v
	}

	return nil
}

func (self *DataSet) IterateData(callback func(int64, interface{}) bool) {
	for k, v := range self.valueMap {
		if !callback(k, v) {
			return
		}
	}
}

func (self *DataSet) SaveData() int {
	var saveCount int

	for _, v := range self.valueMap {

		if vv, ok := v.(IDataSaver); ok && vv.IsDirty() {

			saveCount += vv.SaveData()
		}
	}

	return saveCount
}

func (self *DataSet) IsDirty() bool {
	return true
}

func (self *DataSet) DataCount() int {
	return len(self.valueMap)
}

func (self *DataSet) Clear() {
	self.valueMap = make(map[int64]interface{})
}

func NewDataSet() *DataSet {
	return &DataSet{
		valueMap: make(map[int64]interface{}),
	}
}
