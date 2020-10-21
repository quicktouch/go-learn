package optimizationtest

import "testing"

func TestCrateRequest(t *testing.T) {
	// 将struct转换为字符串
	str := createRequest()
	t.Log(str)
}

func TestProcessRequest(t *testing.T) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	resp := processRequest(reqs)
	t.Log(resp)
}

func BenchmarkName(b *testing.B) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequest(reqs)
	}
	b.StopTimer()
}

/**
➜  optimization git:(master) ✗ go test -bench=. -cpuprofile=cpu.profile
goos: darwin
goarch: amd64
BenchmarkName-8           120354              9887 ns/op
PASS
ok      _/Users/panda/Desktop/go-learn/playground/src/lession/optimization      1.420s
➜  optimization git:(master) ✗ go tool pprof cpu.profile
Type: cpu
Time: Feb 3, 2020 at 12:27pm (CST)
Duration: 1.41s, Total samples = 1.32s (93.76%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 750ms, 56.82% of 1320ms total
Showing top 10 nodes out of 113
      flat  flat%   sum%        cum   cum%
     190ms 14.39% 14.39%      190ms 14.39%  runtime.usleep
     120ms  9.09% 23.48%      180ms 13.64%  strconv.ParseUint
      70ms  5.30% 28.79%       70ms  5.30%  runtime.pthread_cond_signal
      60ms  4.55% 33.33%       60ms  4.55%  github.com/mailru/easyjson/jlexer.(*Lexer).fetchNumber
      60ms  4.55% 37.88%       80ms  6.06%  strconv.FormatInt
      60ms  4.55% 42.42%       60ms  4.55%  strconv.underscoreOK
      50ms  3.79% 46.21%      130ms  9.85%  github.com/mailru/easyjson/jlexer.(*Lexer).FetchToken
      50ms  3.79% 50.00%       50ms  3.79%  runtime.madvise
      50ms  3.79% 53.79%       50ms  3.79%  runtime.pthread_cond_wait
      40ms  3.03% 56.82%      560ms 42.42%  _/Users/panda/Desktop/go-learn/playground/src/lession/optimization.easyjson6a975c40DecodeGoo1
(pprof) top -cum
Showing nodes accounting for 80ms, 6.06% of 1320ms total
Showing top 10 nodes out of 113
      flat  flat%   sum%        cum   cum%
         0     0%     0%      910ms 68.94%  _/Users/panda/Desktop/go-learn/playground/src/lession/optimization.BenchmarkName
      20ms  1.52%  1.52%      910ms 68.94%  _/Users/panda/Desktop/go-learn/playground/src/lession/optimization.processRequest
         0     0%  1.52%      910ms 68.94%  testing.(*B).launch
         0     0%  1.52%      910ms 68.94%  testing.(*B).runN
         0     0%  1.52%      560ms 42.42%  _/Users/panda/Desktop/go-learn/playground/src/lession/optimization.(*Request).UnmarshalJSON
      40ms  3.03%  4.55%      560ms 42.42%  _/Users/panda/Desktop/go-learn/playground/src/lession/optimization.easyjson6a975c40DecodeGoo1
         0     0%  4.55%      330ms 25.00%  runtime.systemstack
         0     0%  4.55%      300ms 22.73%  github.com/mailru/easyjson/jlexer.(*Lexer).Int
      20ms  1.52%  6.06%      300ms 22.73%  github.com/mailru/easyjson/jlexer.(*Lexer).Int64
         0     0%  6.06%      270ms 20.45%  runtime.mstart
(pprof) list processRequest
Total: 1.32s
ROUTINE ======================== _/Users/panda/Desktop/go-learn/playground/src/lession/optimization.processRequest in /Users/panda/Desktop/go-learn/playground/src/lession/optimization/optimization.go
      20ms      910ms (flat, cum) 68.94% of Total
         .          .     22:
         .          .     23:func processRequest(reqs []string) []string {
         .          .     24:   reps := []string{}
         .          .     25:   for _, req := range reqs {
         .          .     26:           reqObj := &Request{}
         .      570ms     27:           reqObj.UnmarshalJSON([]byte(req))
         .          .     28:           //      json.Unmarshal([]byte(req), reqObj)
         .          .     29:
         .          .     30:           var buf strings.Builder
      10ms       10ms     31:           for _, e := range reqObj.PayLoad {
         .      130ms     32:                   buf.WriteString(strconv.Itoa(e))
      10ms       40ms     33:                   buf.WriteString(",")
         .          .     34:           }
         .          .     35:           repObj := &Response{reqObj.TransactionID, buf.String()}
         .      140ms     36:           repJson, err := repObj.MarshalJSON()
         .          .     37:           //repJson, err := json.Marshal(&repObj)
         .          .     38:           if err != nil {
         .          .     39:                   panic(err)
         .          .     40:           }
         .       20ms     41:           reps = append(reps, string(repJson))
         .          .     42:   }
         .          .     43:   return reps
         .          .     44:}
         .          .     45:
         .          .     46:// 接收请求，将其序列化为对象，并处理
*/