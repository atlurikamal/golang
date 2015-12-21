package exp

import (
"net/http"
"fmt"
)
func handler (res http.ResponseWriter, req *http.Request) {
fmt.Fprint(res, "ARE YOU PLANNING FOR YOUR HIGHER STUDIES ABROAD??")
fmt.Fprint(res, "IF YES ENTER /GRE or /GMAT in URL")

switch req.URL.Path{
	
case "/gre": {
	fmt.Fprint(res, "Write GRE exam whose scale start from 160 to 340 \n ")
	fmt.Fprint(res, "Get minimum of 300 score")
	}

case "/gmat":{
	fmt.Fprint(res, "write GMAT exam whose scale start from 200 to 800 \n ")
	fmt.Fprint(res, "get minimum of 550 score")
	}
default:
	fmt.Fprint(res, "nothing is entered in URL")
}

fmt.Fprint(res, "Also write TOEFL or IELTS if you are international student \n ")
fmt.Fprint(res, "Get minimum of score 80 in TOEFL or score 6.5 in IELTS")
}

func init() {
http.HandleFunc("/", handler)
}