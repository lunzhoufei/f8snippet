
"github.com/golang/protobuf/proto"

// SEE: https://findingsea.github.io/2018/09/14/Go-jsonpb/
func readable() {
	var pb abtpb::Experiment
	fmt.Println("experiment info pb.String()=>", pb.String())
	fmt.Println("\n\nexperiment info pb.MarshalTextString()=>", proto.MarshalTextString(&pb))
}

