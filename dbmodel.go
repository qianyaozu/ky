package main


//@工作面
type WorkPlace struct {
	ID        int //工作面编号（主键）
	Name      string        //工作面名称
	DipLength float32       //工作面斜长
	Capacity  float32       //煤层容量
	Length    float32       //回采总长
}

//@钻孔
type Drilling struct {
	ID              int //钻孔编号（主键）
	Name            string        //钻孔名称
	InstallPosition string        //安装位置
	Depth           float32       //钻孔深度
	Distance        float32       //钻孔距离
}

//顶底板位移
type Drift struct {
	ID              int //顶底板位移编号
	Name            string        //顶底板名称
	InstallPosition string        //安装位置
	BrokenDrift     float32       //压并位移
}

//离层
type Displace struct {
	ID              int //离层编号
	InstallPosition string        //安装位置
	Distance        float32       //安装距离
	InstallTime     int32         //安装时间，时间戳
}

//倾角
type Dip struct {
	ID              int //倾角编号
	FrameID         string        //支架编号
	InstallPosition string        //安装位置
}

//支架
type Frame struct {
	ID              int //支架编号
	InstallPosition string
	InitPower       float32 //初撑力
	MaxResistence   float32 //最大工作阻力
}

//支柱
type Pillar struct {
	ID              int //支柱编号
	PillarNo        int           //支柱编号
	PillarType      int           //1单柱 2多柱
	Name            string        //支柱名称
	FrameID         string        //支架编号
	MaxResistence   float32       //最大工作阻力
	InitPower       float32       //支柱初撑力
	DiameterDepth   float32       //缸径深度
	InstallPosition string
}

//基点
type Point struct {
	ID         int //基点编号
	Name       string        //基点名称
	PointType  int           //1单基点，2双基点
	Depth      float32       //基点深度
	InitValue  float32       //初始位移
	DisplaceID string        // 离层编号

}
