package orm

// Setup is 環境構築の最初に行うべきDBのセットアップ
func Setup() {
	dropTable()
	createTable()
}

func dropTable() {
	db.DropTable(&User{}, &FacebookAccount{}, &Skill{}, &SkillUser{}, &Product{}, &Award{}, &License{}, &TwitterAccount{})
}

func createTable() {
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
		if err := insertUser(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&FacebookAccount{}) {
		if err := db.CreateTable(&FacebookAccount{}).Error; err != nil {
			panic(err)
		}
		if err := insertFacebookAccount(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&TwitterAccount{}) {
		if err := db.CreateTable(&TwitterAccount{}).Error; err != nil {
			panic(err)
		}
		if err := insertTwitterAccount(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Skill{}) {
		if err := db.CreateTable(&Skill{}).Error; err != nil {
			panic(err)
		}
		if err := insertSkill(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&SkillUser{}) {
		if err := db.CreateTable(&SkillUser{}).Error; err != nil {
			panic(err)
		}
		if err := insertSkillUser(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Product{}) {
		if err := db.CreateTable(&Product{}).Error; err != nil {
			panic(err)
		}
		if err := insertProduct(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Award{}) {
		if err := db.CreateTable(&Award{}).Error; err != nil {
			panic(err)
		}
		if err := insertAward(); err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&License{}) {
		if err := db.CreateTable(&License{}).Error; err != nil {
			panic(err)
		}
		if err := insertLicense(); err != nil {
			panic(err)
		}
	}
}

func insertUser() (err error) {
	users := []User{
		User{
			Name:              "大河原 拓巳",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/21151682_1723058597990412_7991828349134415504_n.jpg?efg=eyJpIjoidCJ9&oh=d48b585ff9ff2275e19eb8fc706bd144&oe=5AB7058A",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_1.jpg",
			AttributeId:       1,
			Overview:          "事業開発を3度ほど経験し、うち１つを1年で年商3,000万円規模のビジネスに成長させました。主にディレクター、営業としての役割を担っています。最近新たに教育・HR領域でエンジニアと一緒に会社つくろうとしてます。興味のあるエンジニアさん話しましょう！",
			Gender:            1,
			Age:               22,
			Address:           "東京都渋谷区",
			SchoolCarrer:      "東京大学経営学部経営学科(2019年卒業予定)",
			ActivityBase:      "八王子",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "高野 舞",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/17021878_773560269466285_3810550193614137121_n.jpg?oh=81159c07b4743d6324930c6b1cbbab57&oe=5AB737BE",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_2.jpg",
			AttributeId:       2,
			Overview:          "主にサーバーサイドやってます。CAのAbemaTVの部署で長期インターンをしていて、バリバリ働いてます。最近起業にも興味が湧いてきたので、CTOをされているエンジニアさんとかいれば声かけてください。",
			Gender:            2,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学文学部英文学科(2021年卒業予定)",
			ActivityBase:      "代官山",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "木下 美優",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/11904535_462022600637239_3215911181810239589_n.jpg?efg=eyJpIjoidCJ9&oh=e427c996bd6f53d8f9573928787fe225&oe=5AD3BB7B",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_3.jpg",
			AttributeId:       2,
			Overview:          "これまでにベンチャー企業4社でインターンを経験し、そのすべてで優勝した経験があります。今大学２年生なのですが、その経験をいかしてスタートアップ企業で長期インターンをしたいと思っています。どなたかおすすめのスタートアップ企業を紹介してくださると大変嬉しいです。",
			Gender:            2,
			Age:               20,
			Address:           "東京都港区",
			SchoolCarrer:      "東洋大学経営学部経営学科(2020年卒業予定)",
			ActivityBase:      "半蔵門",
			FacebookAccountId: 1,
		},
		User{
			Name:              "水野 慎也",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/22308709_121914455194047_494685754163737983_n.jpg?efg=eyJpIjoidCJ9&oh=a760b9275caed4a26ed967f28860ec57&oe=5AB80091",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_4.jpg",
			AttributeId:       3,
			Overview:          "UI/UXデザインを専門としています。これまで株式会社XYZのUIデザインのインターンで優勝、株式会社サイバーでもインターン優勝経験があります。他にはvivitの頂点デザインコンテストにおいて、グラフィック部門で最優秀賞を受賞しました。ブランディングやパッケージデザインの経験もあります。みなさんとともによりよいサービスを提供できるよう全力を尽くしたいと思います！",
			Gender:            1,
			Age:               22,
			Address:           "東京都新宿区",
			SchoolCarrer:      "多摩美術大学デザイン学科(2019年卒予定)",
			ActivityBase:      "表参道",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "渡辺 拓也",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/15941434_1812752735662130_1421661557137065570_n.jpg?efg=eyJpIjoidCJ9&oh=596a6cf284986e56fa6385671c301640&oe=5ABD7990",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_5.jpg",
			AttributeId:       2,
			Overview:          "大学は文系ですが、現在は休学をしてWebエンジニアとして活動しています。普段は小さい会社で、主にRailsやJavaScript使って開発していて、最近はiOSにも手を出してみたりしてます。ハッカソン型のインターンでは、技術賞やCTO賞といった賞をいただいたことがあります。",
			Gender:            1,
			Age:               20,
			Address:           "東京都品川区",
			SchoolCarrer:      "青山学院大学工学部物理学科(2020年卒業予定)",
			ActivityBase:      "浅草橋",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "山崎 光太郎",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/22282123_1457436660972579_1952528898352429022_n.jpg?efg=eyJpIjoidCJ9&oh=f3db62952c35029a2dac901b37baa765&oe=5AD5A19D",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_6.jpg",
			AttributeId:       1,
			Overview:          "株式会社ジラフにてディレクターとして1年半長期インターンをしています。来年中にSEO関係の事業で起業しようとしています。既に資金調達先も決まっているため、あとはエンジニアさんとデザイナさんで共同創業してくれる人を探すだけです。興味のある方はご連絡ください。",
			Gender:            1,
			Age:               21,
			Address:           "東京都台東区",
			SchoolCarrer:      "筑波大学情報理工学部データサイエンス学科(2018年卒業予定)",
			ActivityBase:      "八王子",
			FacebookAccountId: 1,
		},
		User{
			Name:              "高橋 一生",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/10421318_435400609965646_4768283335432817010_n.jpg?efg=eyJpIjoidCJ9&oh=7c94e2e54d1b5450c88f850239ceb320&oe=5AD3993D",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_7.jpg",
			AttributeId:       1,
			Overview:          "株式会社フラミンゴの元共同創業者です。ビジネスサイドもエンジニアリングもできます。最近は中古スマホのマーケットプレイスサービスや旅行系サービスやってます。どんどん新しい会社をつくっていくスタイルなので、何か事業開発等したい方がいれば気軽に声かけてください。",
			Gender:            1,
			Age:               21,
			Address:           "東京都台東区",
			SchoolCarrer:      "筑波大学情報理工学部データサイエンス学科(2018年卒業予定)",
			ActivityBase:      "恵比寿",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "斎藤 健吾",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/23659417_1987150701563026_2462026853337210000_n.jpg?efg=eyJpIjoidCJ9&oh=903877021710e8265b43d231c584813d&oe=5ACD18A4",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_8.jpg",
			AttributeId:       3,
			Overview:          "最近では商品ブランドを自分で立ち上げ、その商品をPRするためのwebデザインの製作などをしています。ブランドの広告のデザインやパッケージデザインも任されていて、ブランディングとしてのデザインを行なっています。作品を見て興味いただいた方は是非声をかけてください。",
			Gender:            1,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学文学部英文学科(2021年卒業予定)",
			ActivityBase:      "八王子",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "徳永 貴大",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/20046511_1355640164553806_3463604895971062116_n.jpg?efg=eyJpIjoidCJ9&oh=1817aab0965ae39698ee1cff5d1743f6&oe=5AB33901",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_9.jpg",
			AttributeId:       2,
			Overview:          "フロントもサーバーサイドも幅広くできます。最近はSwiftも勉強してiOSアプリをいくつかリリースしました。北海道に住んでいますが、東京に行くこともたくさんあるので、ぜひいろんな人とお話ししてみたいです。",
			Gender:            2,
			Age:               20,
			Address:           "東京都港区",
			SchoolCarrer:      "東洋大学経営学部経営学科(2020年卒業予定)",
			ActivityBase:      "日本橋",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "石井 太一",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/24232528_369422113506760_5937954900653923551_n.jpg?efg=eyJpIjoidCJ9&oh=7ee7751c20393d30c2b7cf562058c099&oe=5ABA17E7",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_10.jpg",
			AttributeId:       3,
			Overview:          "これまではグラフィック系を中心にロゴやパッケージデザインをさせていただくことが多くありました。ブランディングを担当させていただいた食品が無印良品有楽町店で販売されました。作品欄にあるので見て見てください。ブランドのロゴデザインやweb広告などの依頼を待っています。",
			Gender:            1,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学文学部英文学科(2021年卒業予定)",
			ActivityBase:      "大手町",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},

		User{
			Name:              "渡辺 陸斗",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/20476077_1828369570807029_9105758321198584365_n.jpg?efg=eyJpIjoidCJ9&oh=8e140d91dea279e59f434e731aa9f70f&oe=5AB44F2D",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_11.jpg",
			AttributeId:       2,
			Overview:          "コンピュータサークルに所属していて、普段からゲーム、アプリ、サービス色々作ってます。AndroidやUnityが得意で、クライアントサイドを担当することが多いです。あと音声信号処理が専攻です。実績では、9月に行われたCyberAgentさんのKyotoHackで特別賞を頂いたり、大学内のソフトウェアコンテストで優秀賞を受賞したりしました。最近はVR・MRアプリ開発にはまってます。",
			Gender:            1,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学文学部英文学科(2021年卒業予定)",
			ActivityBase:      "日比谷",
			FacebookAccountId: 1,
		},
		User{
			Name:              "中村 太一",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/20840728_1958012951146440_1274856623485082300_n.jpg?efg=eyJpIjoidCJ9&oh=cd861edb98b060016b58817a4bcc9811&oe=5ABB25C6",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_12.jpg",
			AttributeId:       1,
			Overview:          "大学留年しました！笑　時間ができたので、とりあえず起業しようかなと！Hive Shibuyaに毎日通っていて、今はVCの業務を手伝っていますが、アイデアはたくさんあるので、ぜひエンジニアの人お話ししましょう！",
			Gender:            1,
			Age:               20,
			Address:           "東京都品川区",
			SchoolCarrer:      "青山学院大学工学部物理学科(2020年卒業予定)",
			ActivityBase:      "恵比寿",
			FacebookAccountId: 1,
		},
		User{
			Name:              "山田 太郎",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/17353590_1855249878077130_4875411822150264692_n.jpg?efg=eyJpIjoidCJ9&oh=e89f492f77e63d2edc51fc790ec5a012&oe=5AC8A542",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_13.jpg",
			AttributeId:       3,
			Overview:          "デザインと並行してゲームプランニングの勉強をしています。美大生としての感性とプランナーとしてのロジカルシンキングを活かして、より説得性のあるデザインを心掛けています。特に何か作品を作る上で、軸となるコンセプトを考えぬくことを一番に大切にしています。",
			Gender:            1,
			Age:               21,
			Address:           "東京都港区",
			SchoolCarrer:      "東洋大学経営学部経営学科(2020年卒業予定)",
			ActivityBase:      "大手町",
			FacebookAccountId: 1,
		},
		User{
			Name:              "福山 徹",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/408281_191445030988240_1622791291_n.jpg?efg=eyJpIjoidCJ9&oh=dcd571819baed9a1fe7550663daa9d10&oe=5AC24377",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_14.jpg",
			AttributeId:       1,
			Overview:          "学生起業団体BRIGHTNESSで通信機器の販売訪問をしています。200名が所属する組織ですが、去年は売上高ベースで最高記録を樹立しました。この度、東京のスタートアップ企業でこの営業スキルをいかしたいと思っています。どうぞよろしくお願いします。",
			Gender:            1,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学文学部英文学科(2021年卒業予定)",
			ActivityBase:      "日本橋",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "高橋 薫",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/24909892_1930512123935184_6301993005322420762_n.jpg?efg=eyJpIjoidCJ9&oh=562461c91b4f6c03e721b228175e91f2&oe=5AD50759",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_15.jpg",
			AttributeId:       3,
			Overview:          "普段は美大生をしております。時にエンジニア、時にデザイナーと呼ばれる存在です。サービス開発が好きでUI/UXデザインやサービス開発をしています。デザイン・フロント、サーバーサイド何でもやってますが主にrails、時にlaravel・cake です。大学ではProcessing/unityなども書いてます。",
			Gender:            2,
			Age:               21,
			Address:           "東京都台東区",
			SchoolCarrer:      "筑波大学芸術学部デザイン学科(2018年卒業予定)",
			ActivityBase:      "日本橋",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "伊藤 楓",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/23795672_1142473082549670_1079784984423545915_n.jpg?efg=eyJpIjoidCJ9&oh=80d0946f7e285f5413ef75697907622f&oe=5ABCE1F4",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_16.jpg",
			AttributeId:       3,
			Overview:          "グラフィックデザインを専攻していて、趣味でタイポグラフィのデザインやっています。twitterにたくさん作品をあげているので見てみて下さい。ロゴデザインには自信があるので、作品を見ていいと思ってくれた方はどんどん依頼してください！",
			Gender:            1,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学工学部機械工学科(2021年卒業予定)",
			ActivityBase:      "浅草橋",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},

		User{
			Name:              "江渡 美穂",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/20294570_441614052905591_3543768209789804869_n.jpg?efg=eyJpIjoidCJ9&oh=d7d576f3850ef5ec137e27f9e8e8cda5&oe=5A88CDC9",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_17.jpg",
			AttributeId:       1,
			Overview:          "学生起業団体BRIGHTNESSで通信機器の販売訪問をしています。200名が所属する組織ですが、去年は売上高ベースで最高記録を樹立しました。この度、東京のスタートアップ企業でこの営業スキルをいかしたいと思っています。どうぞよろしくお願いします。",
			Gender:            2,
			Age:               21,
			Address:           "東京都台東区",
			SchoolCarrer:      "筑波大学情報理工学部データサイエンス学科(2018年卒業予定)",
			ActivityBase:      "六本木",
			FacebookAccountId: 1,
		},
		User{
			Name:              "太田 聡一",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t31.0-8/fr/cp0/e15/q65/24297451_898904310286288_2469023110063950386_o.jpg?efg=eyJpIjoidCJ9&oh=1c3befc33241dcef0f9830a2b0ff4de0&oe=5AD74DDB",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_18.jpg",
			AttributeId:       3,
			Overview:          "大学ではプロダクトデザインを専攻しています。Rinocerosなどの3DCADの扱いには慣れていて、プロダクトの形状や造形の美しさついて研究してきました。形としてのプロダクトの考案をされている方々の力になれると思います。よろしくお願い致します。",
			Gender:            1,
			Age:               20,
			Address:           "東京都品川区",
			SchoolCarrer:      "青山学院大学工学部物理学科(2020年卒業予定)",
			ActivityBase:      "信濃町",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "鳥居 美仁",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/20246375_827309910769434_2800266932007317777_n.jpg?efg=eyJpIjoidCJ9&oh=004cf924df73c3b9d719aeca8b9f5d11&oe=5AB379AF",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_19.jpg",
			AttributeId:       2,
			Overview:          "名古屋の高専でプログラミングを学んでいます。株式会社エイチームからゲーム開発の仕事で長期インターンを2018年2月から始めることになりました。なので、現在はUnityによるゲーム開発をメインで勉強しています。よろしくお願いします。",
			Gender:            2,
			Age:               20,
			Address:           "東京都品川区",
			SchoolCarrer:      "青山学院大学工学部物理学科(2020年卒業予定)",
			ActivityBase:      "高円寺",
			FacebookAccountId: 1,
		},
		User{
			Name:              "橘 春香",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t31.0-8/fr/cp0/e15/q65/11224032_1187992107883003_3512851942207103721_o.jpg?efg=eyJpIjoidCJ9&oh=508bb28837a7a9ceb9325506c178e18b&oe=5AC2FD82",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_20.jpg",
			AttributeId:       1,
			Overview:          "TABI LABというメディアでライターを２年間続けています。書くことが大好きで、よく海外旅行をしてその旅行記をつづったりもしています。書くお仕事はいつでも募集中ですので、お声がけください。",
			Gender:            2,
			Age:               19,
			Address:           "東京都足立区",
			SchoolCarrer:      "東京大学文学部英文学科(2021年卒業予定)",
			ActivityBase:      "代官山",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "新村 美亜",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/23844960_1585214054898865_667172294994546086_n.jpg?efg=eyJpIjoidCJ9&oh=c94e16e3a4035d4f90541b2afc4b2f90&oe=5ABBD51D",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_21.jpg",
			AttributeId:       1,
			Overview:          "世界一周してきました！東南アジアが大好きで、今はタイのスラム街でe-learning教育を提供しているNPO法人でインターンをしています。優しい世界を実現できたらいいなと。",
			Gender:            2,
			Age:               20,
			Address:           "東京都港区",
			SchoolCarrer:      "東洋大学経営学部経営学科(2020年卒業予定)",
			ActivityBase:      "六本木",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "廣島 海斗",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t31.0-8/fr/cp0/e15/q65/22096132_1340180976092857_642215417203081337_o.jpg?efg=eyJpIjoidCJ9&oh=af208709f4fa1cbaebb0fefc26796b26&oe=5AC34E57",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_22.jpg",
			AttributeId:       2,
			Overview:          "サーバーサイド技術が大好きで、先日行われたCyberAgentさんのアドテクチャレンジで優勝させて頂きました!大規模分散処理に興味があり、パフォーマンスチューニングを自分の武器としたいと考えております。",
			Gender:            1,
			Age:               21,
			Address:           "東京都台東区",
			SchoolCarrer:      "筑波大学情報理工学部データサイエンス学科(2018年卒業予定)",
			ActivityBase:      "六本木",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "伊達 流留",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/19429785_105097533453621_3756725723111449017_n.jpg?efg=eyJpIjoidCJ9&oh=53c12af9e2065c0ced9d8431ad6ef7f0&oe=5ACC5F42",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_23.jpg",
			AttributeId:       1,
			Overview:          "マーケティングが大好きすぎて、あの世界のコトラーの講演にも以前参加しました！マーケティングを概念から戦略に落とし込むプロセスはあまり一般的ではありませんが、今マーケティング4.0という概念から戦略を生み出す必要があると考えています。一緒に戦略練りましょう。",
			Gender:            2,
			Age:               20,
			Address:           "東京都品川区",
			SchoolCarrer:      "青山学院大学工学部物理学科(2020年卒業予定)",
			ActivityBase:      "表参道",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "田中 渉",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/18274946_285521105207242_7792923012535700106_n.jpg?efg=eyJpIjoidCJ9&oh=894c36538c46b82eda1974d4bc67709c&oe=5AB691F3",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_24.jpg",
			AttributeId:       2,
			Overview:          "株式会社Fincにて新規事業サービスを開発していました。ライブストリーミング事業です。この度その知見を生かして起業することになったので、一緒に創業してくれるメンバーを探しています。iOSエンジニアとAndoroidエンジニアの方がいればぜひお声がけください。すでに資金調達するところも決まっているので、かなり面白いと思います！",
			Gender:            1,
			Age:               21,
			Address:           "東京都台東区",
			SchoolCarrer:      "筑波大学情報理工学部データサイエンス学科(2018年卒業予定)",
			ActivityBase:      "表参道",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "大倉 美咲",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-1/25498470_521012341605310_4971587071954879245_n.jpg?oh=f55764c6ec03f5b8b29eb9d5fee9631f&oe=5AB77AF5",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_25.jpg",
			AttributeId:       3,
			Overview:          "制作会社で1年間フロントエンドエンジニアの実務を経て、フリーランスでUI デザイナーをやりつつインターンなどでAIを活用した新規事業やサービス企画、UI設計と横断的なUXデザインを行ってきました。",
			Gender:            2,
			Age:               22,
			Address:           "東京都足立区",
			SchoolCarrer:      "明治大学工学部機械工学科(2019年卒予定)",
			ActivityBase:      "六本木",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "花田 涼介",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/fr/cp0/e15/q65/22045762_694780590715019_1237421679959396786_n.jpg?efg=eyJpIjoidCJ9&oh=b148874576d583e69f878ed39dcac6da&oe=5AC0D65A",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_26.jpg",
			AttributeId:       2,
			Overview:          "去年、リゾート地セブ島にてプログラミングを学ぶという癖のある留学をしていました。",
			Gender:            1,
			Age:               19,
			Address:           "東京都墨田区",
			SchoolCarrer:      "明治大学工学部機械工学科(2019年卒予定)",
			ActivityBase:      "日比谷",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "香田 英二",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/11693907_724445897699721_6359951278343080047_n.jpg?oh=412f7374abfa51947482ad035297cba8&oe=5AC4E326",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_27.jpg",
			AttributeId:       2,
			Overview:          "主にandroid開発(java/kotlin)をしています。現在はスタートアップの企業でandroidアプリ開発のアルバイトをしています。また、今年の夏にはCyberAgentの長期就業型のインターンシップに参加し、アメーバブログのプロダクトでandroidアプリ開発をしていました。学校ではSP2LCという情報系の学生サークルに参加し、高専プログラミングコンテストやセキュリティコンテストに出場するなどの活動を行っています。",
			Gender:            1,
			Age:               19,
			Address:           "東京都港区",
			SchoolCarrer:      "東洋大学経営学部経営学科(2020年卒業予定)",
			ActivityBase:      "六本木",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "越前 龍馬",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t31.0-8/fr/cp0/e15/q65/16601582_2237242236501710_3981390430534250708_o.jpg?efg=eyJpIjoidCJ9&oh=fafa6354227bfd0d4df0d818d3825b57&oe=5AB9A64C",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_28.jpg",
			AttributeId:       3,
			Overview:          "普段は東都大学に通いながら独学でUIデザインを学んでいます。複数企業のUIデザインのインターンに参加し、賞をいただきました。まだ実際にサービスをリリースした経験がないため、いろんな方々とつながってサービス作りを行いたいと思っています。",
			Gender:            1,
			Age:               20,
			Address:           "東京都港区",
			SchoolCarrer:      "東洋大学経営学部経営学科(2020年卒業予定)",
			ActivityBase:      "六本木",
			FacebookAccountId: 1,
			TwitterAccountId:  1,
		},
		User{
			Name:              "大野 加奈",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t1.0-9/13754330_1751694285070218_305240522372902206_n.jpg?oh=b1bb2a2216c10b7525cae25a1143abf9&oe=5AC76B40",
			CoverUrl:          "https://storage.googleapis.com/genepse-186713.appspot.com/cover_images/haikei_29.jpg",
			AttributeId:       3,
			Overview:          "日頃は美大に通っていますが、あるデザイン事務所でバイトをさせていただく中で、大学では学ぶことができないような様々な経験をさせていただきました。UIデザインを得意としていて、実際のサービスリリースの経験もあります。UIデザイナーとして活躍できる場を探しているので、どんどんお声かけください。",
			Gender:            1,
			Age:               19,
			Address:           "東京都品川区",
			SchoolCarrer:      "東京造形大学デザイン学科(2020年卒業予定)",
			ActivityBase:      "信濃町",
			FacebookAccountId: 1,
		},
		User{
			Name:              "吉田 剛",
			AvatarUrl:         "https://scontent-nrt1-1.xx.fbcdn.net/v/t31.0-8/fr/cp0/e15/q65/14543957_1679232499060637_7711106156790730314_o.jpg?efg=eyJpIjoidCJ9&oh=d45a247f575939b5c017b331f078ce8c&oe=5ABD09CA",
			FacebookAccountId: 1,
		},
	}

	for _, u := range users {
		if err = u.Insert(); err != nil {
			return
		}
	}
	return
}

func insertFacebookAccount() (err error) {
	facebooks := []FacebookAccount{
		FacebookAccount{
			AccountId: "",
			MypageUrl: "https://www.facebook.com/yohei.kanatani.5",
		},
	}
	for _, f := range facebooks {
		if err = f.Insert(); err != nil {
			return
		}
	}
	return
}

func insertTwitterAccount() (err error) {
	twitters := []TwitterAccount{
		TwitterAccount{
			AccountId: "willbehulk",
			MypageUrl: "https://twitter.com/willbehulk",
		},
	}
	for _, t := range twitters {
		if err = t.Insert(); err != nil {
			return
		}
	}
	return
}

func insertSkill() (err error) {
	names := []string{
		"事業開発", "投資家", "営業", "法人営業", "経理", "会計", "HR", "法務", "労務", "ライター", "VC", "マーケ", "採用", "R&D", "企画", "Director", "PM", "経営", "起業", "PR", "弁護士", "商品開発", "総務", "秘書", "監査", "税務", "税理士", "品質管理", "財務", "広報", "CEO", "COO", "CFO", "CXO", "CMO",
		"iOS", "Android", "VR", "AR", "Ruby", "Python", "MySQL", "機械学習", "NLP", "Unity", "Java", "PHP", "AWS", "GCP", "Swift", "動画配信", "HTML", "CSS", "JS", "jQuery", "React.js", "Node.js", "CTO",
		"illustrator", "Photoshop", "After Effect", "XD", "Premire", "InDesign", "Sketch", "Prott", "ProtoPie", "Fusion", "Rhinoceros", "Dreamweaver", "Studio", "CINEMA 4D", "Blender", "Maya", "KeyShot", "123D", "ZBrush", "Shade", "Lightwave3D", "V-ray"}

	for _, name := range names {
		s := &Skill{
			Name: name,
		}
		if err = s.Insert(); err != nil {
			return
		}
	}
	return
}

func insertSkillUser() (err error) {
	skillUsers := []SkillUser{
		SkillUser{
			SkillId:   19,
			UserId:    1,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   3,
			UserId:    1,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   12,
			UserId:    1,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   19,
			UserId:    1,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   36,
			UserId:    2,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   46,
			UserId:    2,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   45,
			UserId:    2,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   38,
			UserId:    3,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   45,
			UserId:    3,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   39,
			UserId:    3,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   59,
			UserId:    4,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   65,
			UserId:    4,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   67,
			UserId:    4,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   60,
			UserId:    4,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   62,
			UserId:    4,
			DispOrder: 5,
		},
		SkillUser{
			SkillId:   68,
			UserId:    4,
			DispOrder: 6,
		},
		SkillUser{
			SkillId:   48,
			UserId:    5,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   8,
			UserId:    6,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   9,
			UserId:    6,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   12,
			UserId:    6,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   19,
			UserId:    7,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   29,
			UserId:    7,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   5,
			UserId:    7,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   60,
			UserId:    8,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   61,
			UserId:    8,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   62,
			UserId:    8,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   36,
			UserId:    9,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   50,
			UserId:    9,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   54,
			UserId:    9,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   55,
			UserId:    9,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   52,
			UserId:    9,
			DispOrder: 5,
		},
		SkillUser{
			SkillId:   63,
			UserId:    10,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   61,
			UserId:    10,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   74,
			UserId:    10,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   40,
			UserId:    11,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   47,
			UserId:    11,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   42,
			UserId:    11,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   41,
			UserId:    11,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   49,
			UserId:    11,
			DispOrder: 5,
		},
		SkillUser{
			SkillId:   12,
			UserId:    12,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   3,
			UserId:    12,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   15,
			UserId:    12,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   60,
			UserId:    13,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   62,
			UserId:    13,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   59,
			UserId:    13,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   4,
			UserId:    14,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   15,
			UserId:    14,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   3,
			UserId:    14,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   61,
			UserId:    15,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   66,
			UserId:    15,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   64,
			UserId:    15,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   67,
			UserId:    16,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   65,
			UserId:    16,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   59,
			UserId:    16,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   3,
			UserId:    17,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   12,
			UserId:    17,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   15,
			UserId:    17,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   66,
			UserId:    18,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   61,
			UserId:    18,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   62,
			UserId:    18,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   39,
			UserId:    19,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   45,
			UserId:    19,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   43,
			UserId:    19,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   41,
			UserId:    19,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   48,
			UserId:    19,
			DispOrder: 5,
		},
		SkillUser{
			SkillId:   10,
			UserId:    20,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   15,
			UserId:    20,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   12,
			UserId:    20,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   10,
			UserId:    21,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   3,
			UserId:    21,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   15,
			UserId:    21,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   43,
			UserId:    22,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   40,
			UserId:    22,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   12,
			UserId:    23,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   29,
			UserId:    23,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   3,
			UserId:    23,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   37,
			UserId:    24,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   46,
			UserId:    24,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   54,
			UserId:    24,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   71,
			UserId:    25,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   59,
			UserId:    25,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   68,
			UserId:    25,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   60,
			UserId:    25,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   61,
			UserId:    25,
			DispOrder: 5,
		},
		SkillUser{
			SkillId:   41,
			UserId:    26,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   54,
			UserId:    26,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   44,
			UserId:    26,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   37,
			UserId:    27,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   46,
			UserId:    27,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   36,
			UserId:    27,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   65,
			UserId:    28,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   67,
			UserId:    28,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   64,
			UserId:    28,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   59,
			UserId:    28,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   60,
			UserId:    28,
			DispOrder: 5,
		},
		SkillUser{
			SkillId:   68,
			UserId:    29,
			DispOrder: 1,
		},
		SkillUser{
			SkillId:   75,
			UserId:    29,
			DispOrder: 2,
		},
		SkillUser{
			SkillId:   59,
			UserId:    29,
			DispOrder: 3,
		},
		SkillUser{
			SkillId:   64,
			UserId:    29,
			DispOrder: 4,
		},
		SkillUser{
			SkillId:   62,
			UserId:    29,
			DispOrder: 5,
		},
	}
	for _, s := range skillUsers {
		if err = s.Insert(); err != nil {
			return
		}
	}
	return
}

func insertProduct() (err error) {
	for i := 1; i < 30; i++ {
		product := Product{
			UserId:       uint(i),
			ReferenceUrl: "https://kentaiwami.jp/portfolio/",
			ImageUrl:     "https://storage.googleapis.com/genepse-186713.appspot.com/product_images/Sun Dec 10 06:05:44 UTC 2017.jpg",
			DispOrder:    1,
		}
		if err = product.Insert(); err != nil {
			return
		}
	}

	/*	for _, p := range products {
			if err = p.Insert(); err != nil {
				return
			}
		}
	*/
	return
}

func insertAward() (err error) {
	awards := []Award{
		Award{
			UserId: 1,
			Name:   "ワークスアプリケーションズ入社パス",
		},
		Award{
			UserId: 2,
			Name:   "isucon優勝",
		},
		Award{
			UserId: 2,
			Name:   "AtCoder 2017春季大会準優勝",
		},
		Award{
			UserId: 3,
			Name:   "株式会社Abic 最優秀インターン",
		},
		Award{
			UserId: 4,
			Name:   "GoodDesign賞受賞",
		},
		Award{
			UserId: 4,
			Name:   "PM準優勝",
		},
		Award{
			UserId: 5,
			Name:   "ジロッカソンteratail賞",
		},
		Award{
			UserId: 6,
			Name:   "ジラフ敢闘賞",
		},
		Award{
			UserId: 7,
			Name:   "メルカリBOLDインターン最優秀賞",
		},
		Award{
			UserId: 7,
			Name:   "CAハッカソン優勝",
		},
		Award{
			UserId: 8,
			Name:   "DC優秀者",
		},
		Award{
			UserId: 8,
			Name:   "GoodDesign賞受賞",
		},
		Award{
			UserId: 9,
			Name:   "JPhacks優勝",
		},
		Award{
			UserId: 10,
			Name:   "CAデザイン優勝",
		},
		Award{
			UserId: 10,
			Name:   "GoodDesign賞受賞",
		},
		Award{
			UserId: 11,
			Name:   "Top corder 世界100位",
		},
		Award{
			UserId: 12,
			Name:   "VC Conference 入賞",
		},
		Award{
			UserId: 13,
			Name:   "CAデザイン優勝",
		},
		Award{
			UserId: 13,
			Name:   "GoodDesign賞受賞",
		},
		Award{
			UserId: 18,
			Name:   "メルカリデザイン入賞",
		},
		Award{
			UserId: 18,
			Name:   "未踏エンジニア",
		},
		Award{
			UserId: 20,
			Name:   "DeNAディレクターインターン最優秀賞",
		},
		Award{
			UserId: 20,
			Name:   "メディアConference 最優秀賞",
		},
		Award{
			UserId: 21,
			Name:   "教育マインド賞",
		},
		Award{
			UserId: 21,
			Name:   "Life is Tech 最高メンター賞",
		},
		Award{
			UserId: 21,
			Name:   "BCGインターン最優秀賞",
		},
		Award{
			UserId: 22,
			Name:   "アドテクチャレンジ 優勝",
		},
		Award{
			UserId: 23,
			Name:   "マーケットconference 最優秀賞",
		},
		Award{
			UserId: 24,
			Name:   "ICPC アジア30位",
		},
		Award{
			UserId: 26,
			Name:   "TC登壇",
		},
		Award{
			UserId: 28,
			Name:   "メルカリデザイン賞 入賞",
		},
	}
	for _, a := range awards {
		if err = a.Insert(); err != nil {
			return
		}
	}
	return
}

func insertLicense() (err error) {
	licenses := []License{
		License{
			UserId: 1,
			Name:   "英検1級",
		},
		License{
			UserId: 1,
			Name:   "ダイダラボッチ交友認定証",
		},
		License{
			UserId: 2,
			Name:   "応用情報技術者",
		},
		License{
			UserId: 3,
			Name:   "ネットワークスペシャリスト",
		},
		License{
			UserId: 6,
			Name:   "プロジェクションマッピング師範級",
		},
		License{
			UserId: 6,
			Name:   "スペイン語検定2級",
		},
		License{
			UserId: 8,
			Name:   "イラストレーター検定１級",
		},
		License{
			UserId: 9,
			Name:   "基本情報技術者",
		},
		License{
			UserId: 10,
			Name:   "毛筆5段",
		},
		License{
			UserId: 10,
			Name:   "硬筆８段",
		},
		License{
			UserId: 11,
			Name:   "Ruby Gold",
		},
		License{
			UserId: 12,
			Name:   "ダイダラボッチ交友認定証",
		},
		License{
			UserId: 12,
			Name:   "英検1級",
		},
		License{
			UserId: 13,
			Name:   "硬筆8段",
		},
		License{
			UserId: 13,
			Name:   "フォトショ検定2級",
		},
		License{
			UserId: 14,
			Name:   "スペイン語検定2級",
		},
		License{
			UserId: 14,
			Name:   "プロジェクションマッピング師範級",
		},
		License{
			UserId: 15,
			Name:   "フォトショ検定2級",
		},
		License{
			UserId: 15,
			Name:   "イラストレーター検定1級",
		},
		License{
			UserId: 17,
			Name:   "Googleアナリスト",
		},
		License{
			UserId: 17,
			Name:   "スペイン語検定2級",
		},
		License{
			UserId: 15,
			Name:   "ダイダラポッチ交友認定証",
		},
		License{
			UserId: 19,
			Name:   "データベーススペシャリスト",
		},
		License{
			UserId: 20,
			Name:   "英検1級",
		},
		License{
			UserId: 21,
			Name:   "プロジェクションマッピング師範級",
		},
		License{
			UserId: 22,
			Name:   "LPIC2",
		},
		License{
			UserId: 22,
			Name:   "Ruby Gold",
		},
		License{
			UserId: 23,
			Name:   "Googleアナリスト資格",
		},
		License{
			UserId: 23,
			Name:   "スペイン語検定2級",
		},
		License{
			UserId: 24,
			Name:   "公認会計士",
		},
		License{
			UserId: 24,
			Name:   "基本情報技術者",
		},
		License{
			UserId: 29,
			Name:   "毛筆4段",
		},
	}
	for _, l := range licenses {
		if err = l.Insert(); err != nil {
			return
		}
	}
	return
}
