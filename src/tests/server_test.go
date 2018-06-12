package tests

//import "quarto/serializer"
//import "quarto/server"
import "quarto/game"
import "testing"
//import "reflect"
//import "net/http"

func TestEndPointForSuggestMove(t *testing.T) {
	//var w http.ResponseWriter
	//var r *http.Request
	
	referenceState := game.GetNewState(4)
	referenceState.Grid[0] = []int{4,5,1,3}
	referenceState.Grid[2] = []int{9,7,10,0}
	referenceState.Piece = 2

	//server.SuggestMove(w, r)
	//b :=[]byte("{\"Grid\":[[4,5,1,3],[0,0,0,0],[9,7,10,0],[0,0,0,0]],\"Piece\":2}")
	//var state, err = serializer.FromJSONToState(b)
	
	//if err != nil {
	//	t.Error("Unserializing shouldn't raise an error")
	//}
}
