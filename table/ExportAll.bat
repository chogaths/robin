REM 电子表格文件名=表名=类型名+File
mkdir obj
call xlsm2pbt.bat Activity
call xlsm2pbt.bat Actor
call xlsm2pbt.bat ActorEvent
call xlsm2pbt.bat Buff
call xlsm2pbt.bat Event
call xlsm2pbt.bat Item
call xlsm2pbt.bat Skill
call xlsm2pbt.bat Level
call xlsm2pbt.bat Quest
call xlsm2pbt.bat HeroExp
call xlsm2pbt.bat MercExp
call xlsm2pbt.bat MercStarExp
call xlsm2pbt.bat AI
call xlsm2pbt.bat Package Package "-sheet=Package;Roulette;RunnePropertyRank;Refine" 
call xlsm2pbt.bat Composite
call xlsm2pbt.bat Fetter
call xlsm2pbt.bat Numerical Numerical "-sheet=EquipStrengthen;EquipAdvance;EquipRefine;EquipRune;LeadHeroUpgrade;MercUpgrade;MercStar;ActorProperty" 
call xlsm2pbt.bat Global
call xlsm2pbt.bat ItemRank
call xlsm2pbt.bat I18n
call xlsm2pbt.bat Name
call xlsm2pbt.bat Goods
call xlsm2pbt.bat MystGoods Goods
call xlsm2pbt.bat VIP
call xlsm2pbt.bat Effect
call xlsm2pbt.bat Speciality
call xlsm2pbt.bat OnlinePackage
call xlsm2pbt.bat NewGuideOrder
call xlsm2pbt.bat Guide
call xlsm2pbt.bat Audio
call xlsm2pbt.bat EquipStren
call xlsm2pbt.bat Tips
call xlsm2pbt.bat BagUnlock Item
call xlsm2pbt.bat RuneSkill
call xlsm2pbt.bat Plot
call xlsm2pbt.bat PlotEnd "-sheet=EquipStrengthen;EquipAdvance;EquipRefine;EquipRune;LeadHeroUpgrade;MercUpgrade;MercStar;ActorProperty" 

