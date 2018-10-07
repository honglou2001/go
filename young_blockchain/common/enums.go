package common

type State int

// iota after initialed, the value will automate increase
const (
	Miner State = iota   // value --> 0
	Transaction          // value --> 1
	Listaddresses        // value --> 2
	Printchain           // value --> 3
	Getbalance			 // value --> 4
)

func (this State) String() string {
	switch this {
	case Miner:
		return "Miner"
	case Transaction:
		return "Transaction"
	case Listaddresses:
		return "Listaddresses"
	case Printchain:
		return "Printchain"
	case Getbalance:
		return "Getbalance"
	default:
		return "Unknow"
	}
}


//func main() {
//	state := miner
//	fmt.Println("state", state)
//}
// output state Running
// if not override String,the state output would be 0
