package hydraDBLayer

// func BenchmarkMySQLDBReads(b *testing.B) {
// 	db.layer, err := ConnectDatabase("mysql", "gouser:gouser@/Hydra")
// 	if err != nil {
// 		b.Fatal("could not connect to hydra chat system",err)
// 	}
// 	findMembersBM(b,dblayer)
// }

// func BenchmarkMongoDBReads(b *testing.B) {
// 	dblayer, err := ConnectDatabase("mongodb", "mongodb://localhost")
// 	if err != nil {
// 		b.Error("Could not connect to hydra chat system", err)
// 		return
// 	}
// }

// func findMembersBM(b *testing.B, dblayer DBLayer) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < b.N; i++ {
// 		_, err := dblayer.FindMember(rand.Intn(16) + 1)
// 		if err != nil {
// 			b.Error("Query failed", err)
// 			return
// 		}
// 	}
// }
