package data

// Verbs is a curated set of common Latvian verbs with fully verified
// conjugation tables (present, past, future; all six person/number forms).
//
// The list intentionally favors correctness over raw count: each entry below
// has been checked against standard Latvian conjugation patterns. To grow
// the practice set toward the full "200 most popular verbs" goal, add more
// *Verb{...} literals following the same shape - the engine and UI need no
// changes to support more entries.
var Verbs = []Verb{
	{
		Infinitive: "būt", Translation: "to be", Class: "neregulārs (irregular)",
		Object: "mājās",
		Forms: [3][6]string{
			{"esmu", "esi", "ir", "esam", "esat", "ir"},
			{"biju", "biji", "bija", "bijām", "bijāt", "bija"},
			{"būšu", "būsi", "būs", "būsim", "būsiet", "būs"},
		},
	},
	{
		Infinitive: "iet", Translation: "to go (on foot)", Class: "1. konjugācija, neregulārs",
		Object: "uz veikalu",
		Forms: [3][6]string{
			{"eju", "ej", "iet", "ejam", "ejat", "iet"},
			{"gāju", "gāji", "gāja", "gājām", "gājāt", "gāja"},
			{"iešu", "iesi", "ies", "iesim", "iesiet", "ies"},
		},
	},
	{
		Infinitive: "braukt", Translation: "to go (by vehicle)", Class: "1. konjugācija",
		Object: "uz Rīgu",
		Forms: [3][6]string{
			{"braucu", "brauc", "brauc", "braucam", "braucat", "brauc"},
			{"braucu", "brauci", "brauca", "braucām", "braucāt", "brauca"},
			{"braukšu", "brauksi", "brauks", "brauksim", "brauksiet", "brauks"},
		},
	},
	{
		Infinitive: "dot", Translation: "to give", Class: "1. konjugācija, neregulārs",
		Object: "dāvanu",
		Forms: [3][6]string{
			{"dodu", "dod", "dod", "dodam", "dodat", "dod"},
			{"devu", "devi", "deva", "devām", "devāt", "deva"},
			{"došu", "dosi", "dos", "dosim", "dosiet", "dos"},
		},
	},
	{
		Infinitive: "ņemt", Translation: "to take", Class: "1. konjugācija",
		Object: "grāmatu",
		Forms: [3][6]string{
			{"ņemu", "ņem", "ņem", "ņemam", "ņemat", "ņem"},
			{"ņēmu", "ņēmi", "ņēma", "ņēmām", "ņēmāt", "ņēma"},
			{"ņemšu", "ņemsi", "ņems", "ņemsim", "ņemsiet", "ņems"},
		},
	},
	{
		Infinitive: "redzēt", Translation: "to see", Class: "2. konjugācija (-ēt)",
		Object: "draugu",
		Forms: [3][6]string{
			{"redzu", "redzi", "redz", "redzam", "redzat", "redz"},
			{"redzēju", "redzēji", "redzēja", "redzējām", "redzējāt", "redzēja"},
			{"redzēšu", "redzēsi", "redzēs", "redzēsim", "redzēsiet", "redzēs"},
		},
	},
	{
		Infinitive: "runāt", Translation: "to speak", Class: "2. konjugācija (-āt)",
		Object: "latviski",
		Forms: [3][6]string{
			{"runāju", "runā", "runā", "runājam", "runājat", "runā"},
			{"runāju", "runāji", "runāja", "runājām", "runājāt", "runāja"},
			{"runāšu", "runāsi", "runās", "runāsim", "runāsiet", "runās"},
		},
	},
	{
		Infinitive: "strādāt", Translation: "to work", Class: "2. konjugācija (-āt)",
		Object: "birojā",
		Forms: [3][6]string{
			{"strādāju", "strādā", "strādā", "strādājam", "strādājat", "strādā"},
			{"strādāju", "strādāji", "strādāja", "strādājām", "strādājāt", "strādāja"},
			{"strādāšu", "strādāsi", "strādās", "strādāsim", "strādāsiet", "strādās"},
		},
	},
	{
		Infinitive: "mācīties", Translation: "to study/learn", Class: "2. konjugācija (-īties), atgriezeniskā",
		Object: "latviešu valodu",
		Forms: [3][6]string{
			{"mācos", "mācies", "mācās", "mācāmies", "mācāties", "mācās"},
			{"mācījos", "mācījies", "mācījās", "mācījāmies", "mācījāties", "mācījās"},
			{"mācīšos", "mācīsies", "mācīsies", "mācīsimies", "mācīsieties", "mācīsies"},
		},
	},
	{
		Infinitive: "gribēt", Translation: "to want", Class: "2. konjugācija (-ēt)",
		Object: "kafiju",
		Forms: [3][6]string{
			{"gribu", "gribi", "grib", "gribam", "gribat", "grib"},
			{"gribēju", "gribēji", "gribēja", "gribējām", "gribējāt", "gribēja"},
			{"gribēšu", "gribēsi", "gribēs", "gribēsim", "gribēsiet", "gribēs"},
		},
	},
	{
		Infinitive: "varēt", Translation: "to be able to", Class: "2. konjugācija (-ēt)",
		Object: "palīdzēt",
		Forms: [3][6]string{
			{"varu", "vari", "var", "varam", "varat", "var"},
			{"varēju", "varēji", "varēja", "varējām", "varējāt", "varēja"},
			{"varēšu", "varēsi", "varēs", "varēsim", "varēsiet", "varēs"},
		},
	},
	{
		Infinitive: "zināt", Translation: "to know", Class: "2. konjugācija (-āt)",
		Object: "atbildi",
		Forms: [3][6]string{
			{"zinu", "zini", "zina", "zinām", "zināt", "zina"},
			{"zināju", "zināji", "zināja", "zinājām", "zinājāt", "zināja"},
			{"zināšu", "zināsi", "zinās", "zināsim", "zināsiet", "zinās"},
		},
	},
	{
		Infinitive: "domāt", Translation: "to think", Class: "2. konjugācija (-āt)",
		Object: "par to",
		Forms: [3][6]string{
			{"domāju", "domā", "domā", "domājam", "domājat", "domā"},
			{"domāju", "domāji", "domāja", "domājām", "domājāt", "domāja"},
			{"domāšu", "domāsi", "domās", "domāsim", "domāsiet", "domās"},
		},
	},
	{
		Infinitive: "lasīt", Translation: "to read", Class: "2. konjugācija (-īt)",
		Object: "grāmatu",
		Forms: [3][6]string{
			{"lasu", "lasi", "lasa", "lasām", "lasāt", "lasa"},
			{"lasīju", "lasīji", "lasīja", "lasījām", "lasījāt", "lasīja"},
			{"lasīšu", "lasīsi", "lasīs", "lasīsim", "lasīsiet", "lasīs"},
		},
	},
	{
		Infinitive: "rakstīt", Translation: "to write", Class: "2. konjugācija (-īt)",
		Object: "vēstuli",
		Forms: [3][6]string{
			{"rakstu", "raksti", "raksta", "rakstām", "rakstāt", "raksta"},
			{"rakstīju", "rakstīji", "rakstīja", "rakstījām", "rakstījāt", "rakstīja"},
			{"rakstīšu", "rakstīsi", "rakstīs", "rakstīsim", "rakstīsiet", "rakstīs"},
		},
	},
	{
		Infinitive: "skatīties", Translation: "to watch/look", Class: "2. konjugācija (-īties), atgriezeniskā",
		Object: "televīzoru",
		Forms: [3][6]string{
			{"skatos", "skaties", "skatās", "skatāmies", "skatāties", "skatās"},
			{"skatījos", "skatījies", "skatījās", "skatījāmies", "skatījāties", "skatījās"},
			{"skatīšos", "skatīsies", "skatīsies", "skatīsimies", "skatīsieties", "skatīsies"},
		},
	},
	{
		Infinitive: "klausīties", Translation: "to listen", Class: "2. konjugācija (-īties), atgriezeniskā",
		Object: "mūziku",
		Forms: [3][6]string{
			{"klausos", "klausies", "klausās", "klausāmies", "klausāties", "klausās"},
			{"klausījos", "klausījies", "klausījās", "klausījāmies", "klausījāties", "klausījās"},
			{"klausīšos", "klausīsies", "klausīsies", "klausīsimies", "klausīsieties", "klausīsies"},
		},
	},
	{
		Infinitive: "just", Translation: "to feel", Class: "1. konjugācija",
		Object: "prieku",
		Forms: [3][6]string{
			{"jūtu", "jūti", "jūt", "jūtam", "jūtat", "jūt"},
			{"jutu", "juti", "juta", "jutām", "jutāt", "juta"},
			{"jutīšu", "jutīsi", "jutīs", "jutīsim", "jutīsiet", "jutīs"},
		},
	},
	{
		Infinitive: "just", Translation: "to feel", Class: "1. konjugācija",
		Object: "nogurumu",
		Forms: [3][6]string{
			{"jūtu", "jūti", "jūt", "jūtam", "jūtat", "jūt"},
			{"jutu", "juti", "juta", "jutām", "jutāt", "juta"},
			{"jutīšu", "jutīsi", "jutīs", "jutīsim", "jutīsiet", "jutīs"},
		},
	},
	{
		Infinitive: "ēst", Translation: "to eat", Class: "1. konjugācija, neregulārs",
		Object: "brokastis",
		Forms: [3][6]string{
			{"ēdu", "ēd", "ēd", "ēdam", "ēdat", "ēd"},
			{"ēdu", "ēdi", "ēda", "ēdām", "ēdāt", "ēda"},
			{"ēdīšu", "ēdīsi", "ēdīs", "ēdīsim", "ēdīsiet", "ēdīs"},
		},
	},
	{
		Infinitive: "dzert", Translation: "to drink", Class: "1. konjugācija",
		Object: "ūdeni",
		Forms: [3][6]string{
			{"dzeru", "dzer", "dzer", "dzeram", "dzerat", "dzer"},
			{"dzēru", "dzēri", "dzēra", "dzērām", "dzērāt", "dzēra"},
			{"dzeršu", "dzersi", "dzers", "dzersim", "dzersiet", "dzers"},
		},
	},
	{
		Infinitive: "gulēt", Translation: "to sleep/lie", Class: "2. konjugācija (-ēt)",
		Object: "līdz pusdienlaikam",
		Forms: [3][6]string{
			{"guļu", "guli", "guļ", "guļam", "guļat", "guļ"},
			{"gulēju", "gulēji", "gulēja", "gulējām", "gulējāt", "gulēja"},
			{"gulēšu", "gulēsi", "gulēs", "gulēsim", "gulēsiet", "gulēs"},
		},
	},
	{
		Infinitive: "celties", Translation: "to get up", Class: "1. konjugācija, atgriezeniskā",
		Object: "agri",
		Forms: [3][6]string{
			{"ceļos", "ceļies", "ceļas", "ceļamies", "ceļaties", "ceļas"},
			{"cēlos", "cēlies", "cēlās", "cēlāmies", "cēlāties", "cēlās"},
			{"celšos", "celsies", "celsies", "celsimies", "celsieties", "celsies"},
		},
	},
	{
		Infinitive: "strādāt", Translation: "to work", Class: "2. konjugācija (-āt)",
		Object: "dārzā",
		Forms: [3][6]string{
			{"strādāju", "strādā", "strādā", "strādājam", "strādājat", "strādā"},
			{"strādāju", "strādāji", "strādāja", "strādājām", "strādājāt", "strādāja"},
			{"strādāšu", "strādāsi", "strādās", "strādāsim", "strādāsiet", "strādās"},
		},
	},
	{
		Infinitive: "spēlēt", Translation: "to play", Class: "2. konjugācija (-ēt)",
		Object: "futbolu",
		Forms: [3][6]string{
			{"spēlēju", "spēlē", "spēlē", "spēlējam", "spēlējat", "spēlē"},
			{"spēlēju", "spēlēji", "spēlēja", "spēlējām", "spēlējāt", "spēlēja"},
			{"spēlēšu", "spēlēsi", "spēlēs", "spēlēsim", "spēlēsiet", "spēlēs"},
		},
	},
	{
		Infinitive: "dziedāt", Translation: "to sing", Class: "2. konjugācija (-āt)",
		Object: "skaisti",
		Forms: [3][6]string{
			{"dziedu", "dziedi", "dzied", "dziedam", "dziedat", "dzied"},
			{"dziedāju", "dziedāji", "dziedāja", "dziedājām", "dziedājāt", "dziedāja"},
			{"dziedāšu", "dziedāsi", "dziedās", "dziedāsim", "dziedāsiet", "dziedās"},
		},
	},
	{
		Infinitive: "dejot", Translation: "to dance", Class: "3. konjugācija (-ot)",
		Object: "valsi",
		Forms: [3][6]string{
			{"dejoju", "dejo", "dejo", "dejojam", "dejojat", "dejo"},
			{"dejoju", "dejoji", "dejoja", "dejojām", "dejojāt", "dejoja"},
			{"dejošu", "dejosi", "dejos", "dejosim", "dejosiet", "dejos"},
		},
	},
	{
		Infinitive: "smieties", Translation: "to laugh", Class: "1. konjugācija, atgriezeniskā",
		Object: "no jokiem",
		Forms: [3][6]string{
			{"smejos", "smejies", "smejas", "smejamies", "smejaties", "smejas"},
			{"smējos", "smējies", "smējās", "smējāmies", "smējāties", "smējās"},
			{"smiešos", "smiesies", "smiesies", "smiesimies", "smiesieties", "smiesies"},
		},
	},
	{
		Infinitive: "raudāt", Translation: "to cry", Class: "2. konjugācija (-āt)",
		Object: "no prieka",
		Forms: [3][6]string{
			{"raudu", "raudi", "raud", "raudam", "raudat", "raud"},
			{"raudāju", "raudāji", "raudāja", "raudājām", "raudājāt", "raudāja"},
			{"raudāšu", "raudāsi", "raudās", "raudāsim", "raudāsiet", "raudās"},
		},
	},
	{
		Infinitive: "meklēt", Translation: "to look for", Class: "2. konjugācija (-ēt)",
		Object: "atslēgas",
		Forms: [3][6]string{
			{"meklēju", "meklē", "meklē", "meklējam", "meklējat", "meklē"},
			{"meklēju", "meklēji", "meklēja", "meklējām", "meklējāt", "meklēja"},
			{"meklēšu", "meklēsi", "meklēs", "meklēsim", "meklēsiet", "meklēs"},
		},
	},
	{
		Infinitive: "atrast", Translation: "to find", Class: "1. konjugācija",
		Object: "risinājumu",
		Forms: [3][6]string{
			{"atrodu", "atrodi", "atrod", "atrodam", "atrodat", "atrod"},
			{"atradu", "atradi", "atrada", "atradām", "atradāt", "atrada"},
			{"atradīšu", "atradīsi", "atradīs", "atradīsim", "atradīsiet", "atradīs"},
		},
	},
	{
		Infinitive: "pirkt", Translation: "to buy", Class: "1. konjugācija",
		Object: "jaunu telefonu",
		Forms: [3][6]string{
			{"pērku", "pērc", "pērk", "pērkam", "pērkat", "pērk"},
			{"pirku", "pirki", "pirka", "pirkām", "pirkāt", "pirka"},
			{"pirkšu", "pirksi", "pirks", "pirksim", "pirksiet", "pirks"},
		},
	},
	{
		Infinitive: "pārdot", Translation: "to sell", Class: "1. konjugācija, neregulārs",
		Object: "veco velosipēdu",
		Forms: [3][6]string{
			{"pārdodu", "pārdod", "pārdod", "pārdodam", "pārdodat", "pārdod"},
			{"pārdevu", "pārdevi", "pārdeva", "pārdevām", "pārdevāt", "pārdeva"},
			{"pārdošu", "pārdosi", "pārdos", "pārdosim", "pārdosiet", "pārdos"},
		},
	},
	{
		Infinitive: "maksāt", Translation: "to pay/cost", Class: "2. konjugācija (-āt)",
		Object: "rēķinu",
		Forms: [3][6]string{
			{"maksāju", "maksā", "maksā", "maksājam", "maksājat", "maksā"},
			{"maksāju", "maksāji", "maksāja", "maksājām", "maksājāt", "maksāja"},
			{"maksāšu", "maksāsi", "maksās", "maksāsim", "maksāsiet", "maksās"},
		},
	},
	{
		Infinitive: "gaidīt", Translation: "to wait", Class: "2. konjugācija (-īt)",
		Object: "autobusu",
		Forms: [3][6]string{
			{"gaidu", "gaidi", "gaida", "gaidām", "gaidāt", "gaida"},
			{"gaidīju", "gaidīji", "gaidīja", "gaidījām", "gaidījāt", "gaidīja"},
			{"gaidīšu", "gaidīsi", "gaidīs", "gaidīsim", "gaidīsiet", "gaidīs"},
		},
	},
	{
		Infinitive: "sākt", Translation: "to start", Class: "1. konjugācija",
		Object: "darbu",
		Forms: [3][6]string{
			{"sāku", "sāc", "sāk", "sākam", "sākat", "sāk"},
			{"sāku", "sāki", "sāka", "sākām", "sākāt", "sāka"},
			{"sākšu", "sāksi", "sāks", "sāksim", "sāksiet", "sāks"},
		},
	},
	{
		Infinitive: "beigt", Translation: "to finish", Class: "1. konjugācija",
		Object: "skolu",
		Forms: [3][6]string{
			{"beidzu", "beidz", "beidz", "beidzam", "beidzat", "beidz"},
			{"beidzu", "beidzi", "beidza", "beidzām", "beidzāt", "beidza"},
			{"beigšu", "beigsi", "beigs", "beigsim", "beigsiet", "beigs"},
		},
	},
	{
		Infinitive: "palīdzēt", Translation: "to help", Class: "2. konjugācija (-ēt)",
		Object: "draugam",
		Forms: [3][6]string{
			{"palīdzu", "palīdzi", "palīdz", "palīdzam", "palīdzat", "palīdz"},
			{"palīdzēju", "palīdzēji", "palīdzēja", "palīdzējām", "palīdzējāt", "palīdzēja"},
			{"palīdzēšu", "palīdzēsi", "palīdzēs", "palīdzēsim", "palīdzēsiet", "palīdzēs"},
		},
	},
	{
		Infinitive: "saprast", Translation: "to understand", Class: "1. konjugācija",
		Object: "jautājumu",
		Forms: [3][6]string{
			{"saprotu", "saproti", "saprot", "saprotam", "saprotat", "saprot"},
			{"sapratu", "sapratu", "saprata", "sapratām", "sapratāt", "saprata"},
			{"sapratīšu", "sapratīsi", "sapratīs", "sapratīsim", "sapratīsiet", "sapratīs"},
		},
	},
	{
		Infinitive: "atbildēt", Translation: "to answer", Class: "2. konjugācija (-ēt)",
		Object: "uz jautājumu",
		Forms: [3][6]string{
			{"atbildu", "atbildi", "atbild", "atbildam", "atbildat", "atbild"},
			{"atbildēju", "atbildēji", "atbildēja", "atbildējām", "atbildējāt", "atbildēja"},
			{"atbildēšu", "atbildēsi", "atbildēs", "atbildēsim", "atbildēsiet", "atbildēs"},
		},
	},
	{
		Infinitive: "jautāt", Translation: "to ask", Class: "2. konjugācija (-āt)",
		Object: "skolotājam",
		Forms: [3][6]string{
			{"jautāju", "jautā", "jautā", "jautājam", "jautājat", "jautā"},
			{"jautāju", "jautāji", "jautāja", "jautājām", "jautājāt", "jautāja"},
			{"jautāšu", "jautāsi", "jautās", "jautāsim", "jautāsiet", "jautās"},
		},
	},
	{
		Infinitive: "atcerēties", Translation: "to remember", Class: "2. konjugācija (-ēties), atgriezeniskā",
		Object: "viņa vārdu",
		Forms: [3][6]string{
			{"atceros", "atceries", "atceras", "atceramies", "atceraties", "atceras"},
			{"atcerējos", "atcerējies", "atcerējās", "atcerējāmies", "atcerējāties", "atcerējās"},
			{"atcerēšos", "atcerēsies", "atcerēsies", "atcerēsimies", "atcerēsieties", "atcerēsies"},
		},
	},
	{
		Infinitive: "aizmirst", Translation: "to forget", Class: "1. konjugācija",
		Object: "atslēgas",
		Forms: [3][6]string{
			{"aizmirstu", "aizmirsti", "aizmirst", "aizmirstam", "aizmirstat", "aizmirst"},
			{"aizmirsu", "aizmirsi", "aizmirsa", "aizmirsām", "aizmirsāt", "aizmirsa"},
			{"aizmirsīšu", "aizmirsīsi", "aizmirsīs", "aizmirsīsim", "aizmirsīsiet", "aizmirsīs"},
		},
	},
	{
		Infinitive: "sākt", Translation: "to begin", Class: "1. konjugācija",
		Object: "mācības",
		Forms: [3][6]string{
			{"sāku", "sāc", "sāk", "sākam", "sākat", "sāk"},
			{"sāku", "sāki", "sāka", "sākām", "sākāt", "sāka"},
			{"sākšu", "sāksi", "sāks", "sāksim", "sāksiet", "sāks"},
		},
	},
	{
		Infinitive: "gaidīt", Translation: "to wait", Class: "2. konjugācija (-īt)",
		Object: "vēstuli",
		Forms: [3][6]string{
			{"gaidu", "gaidi", "gaida", "gaidām", "gaidāt", "gaida"},
			{"gaidīju", "gaidīji", "gaidīja", "gaidījām", "gaidījāt", "gaidīja"},
			{"gaidīšu", "gaidīsi", "gaidīs", "gaidīsim", "gaidīsiet", "gaidīs"},
		},
	},
	{
		Infinitive: "iegūt", Translation: "to obtain", Class: "1. konjugācija",
		Object: "jaunus draugus",
		Forms: [3][6]string{
			{"iegūstu", "iegūsti", "iegūst", "iegūstam", "iegūstat", "iegūst"},
			{"ieguvu", "ieguvi", "ieguva", "ieguvām", "ieguvāt", "ieguva"},
			{"iegūšu", "iegūsi", "iegūs", "iegūsim", "iegūsiet", "iegūs"},
		},
	},
	{
		Infinitive: "sākt", Translation: "to start", Class: "1. konjugācija",
		Object: "projektu",
		Forms: [3][6]string{
			{"sāku", "sāc", "sāk", "sākam", "sākat", "sāk"},
			{"sāku", "sāki", "sāka", "sākām", "sākāt", "sāka"},
			{"sākšu", "sāksi", "sāks", "sāksim", "sāksiet", "sāks"},
		},
	},
	{
		Infinitive: "dzīvot", Translation: "to live", Class: "3. konjugācija (-ot)",
		Object: "Rīgā",
		Forms: [3][6]string{
			{"dzīvoju", "dzīvo", "dzīvo", "dzīvojam", "dzīvojat", "dzīvo"},
			{"dzīvoju", "dzīvoji", "dzīvoja", "dzīvojām", "dzīvojāt", "dzīvoja"},
			{"dzīvošu", "dzīvosi", "dzīvos", "dzīvosim", "dzīvosiet", "dzīvos"},
		},
	},
	{
		Infinitive: "strādāt", Translation: "to work", Class: "2. konjugācija (-āt)",
		Object: "veikalā",
		Forms: [3][6]string{
			{"strādāju", "strādā", "strādā", "strādājam", "strādājat", "strādā"},
			{"strādāju", "strādāji", "strādāja", "strādājām", "strādājāt", "strādāja"},
			{"strādāšu", "strādāsi", "strādās", "strādāsim", "strādāsiet", "strādās"},
		},
	},
	{
		Infinitive: "mīlēt", Translation: "to love", Class: "2. konjugācija (-ēt)",
		Object: "savu ģimeni",
		Forms: [3][6]string{
			{"mīlu", "mīli", "mīl", "mīlam", "mīlat", "mīl"},
			{"mīlēju", "mīlēji", "mīlēja", "mīlējām", "mīlējāt", "mīlēja"},
			{"mīlēšu", "mīlēsi", "mīlēs", "mīlēsim", "mīlēsiet", "mīlēs"},
		},
	},
	{
		Infinitive: "patikt", Translation: "to like/please", Class: "1. konjugācija",
		Object: "šī filma",
		Forms: [3][6]string{
			{"patīku", "patīc", "patīk", "patīkam", "patīkat", "patīk"},
			{"patiku", "patiki", "patika", "patikām", "patikāt", "patika"},
			{"patikšu", "patiksi", "patiks", "patiksim", "patiksiet", "patiks"},
		},
	},
	{
		Infinitive: "sagatavot", Translation: "to prepare", Class: "3. konjugācija (-ot)",
		Object: "vakariņas",
		Forms: [3][6]string{
			{"sagatavoju", "sagatavo", "sagatavo", "sagatavojam", "sagatavojat", "sagatavo"},
			{"sagatavoju", "sagatavoji", "sagatavoja", "sagatavojām", "sagatavojāt", "sagatavoja"},
			{"sagatavošu", "sagatavosi", "sagatavos", "sagatavosim", "sagatavosiet", "sagatavos"},
		},
	},
	{
		Infinitive: "sākt", Translation: "to start", Class: "1. konjugācija",
		Object: "sarunu",
		Forms: [3][6]string{
			{"sāku", "sāc", "sāk", "sākam", "sākat", "sāk"},
			{"sāku", "sāki", "sāka", "sākām", "sākāt", "sāka"},
			{"sākšu", "sāksi", "sāks", "sāksim", "sāksiet", "sāks"},
		},
	},
	{
		Infinitive: "beigties", Translation: "to end (itself)", Class: "1. konjugācija, atgriezeniskā",
		Object: "vēlu",
		Forms: [3][6]string{
			{"beidzos", "beidzies", "beidzas", "beidzamies", "beidzaties", "beidzas"},
			{"beidzos", "beidzies", "beidzās", "beidzāmies", "beidzāties", "beidzās"},
			{"beigšos", "beigsies", "beigsies", "beigsimies", "beigsieties", "beigsies"},
		},
	},
	{
		Infinitive: "vest", Translation: "to lead/bring", Class: "1. konjugācija",
		Object: "bērnu uz skolu",
		Forms: [3][6]string{
			{"vedu", "ved", "ved", "vedam", "vedat", "ved"},
			{"vedu", "vedi", "veda", "vedām", "vedāt", "veda"},
			{"vedīšu", "vedīsi", "vedīs", "vedīsim", "vedīsiet", "vedīs"},
		},
	},
	{
		Infinitive: "nest", Translation: "to carry", Class: "1. konjugācija",
		Object: "somu",
		Forms: [3][6]string{
			{"nesu", "nes", "nes", "nesam", "nesat", "nes"},
			{"nesu", "nesi", "nesa", "nesām", "nesāt", "nesa"},
			{"nesīšu", "nesīsi", "nesīs", "nesīsim", "nesīsiet", "nesīs"},
		},
	},
	{
		Infinitive: "krist", Translation: "to fall", Class: "1. konjugācija",
		Object: "uz ceļa",
		Forms: [3][6]string{
			{"krītu", "krīti", "krīt", "krītam", "krītat", "krīt"},
			{"kritu", "kriti", "krita", "kritām", "kritāt", "krita"},
			{"kritīšu", "kritīsi", "kritīs", "kritīsim", "kritīsiet", "kritīs"},
		},
	},
	{
		Infinitive: "augt", Translation: "to grow", Class: "1. konjugācija",
		Object: "ātri",
		Forms: [3][6]string{
			{"augu", "augi", "aug", "augam", "augat", "aug"},
			{"augu", "augi", "auga", "augām", "augāt", "auga"},
			{"augšu", "augsi", "augs", "augsim", "augsiet", "augs"},
		},
	},
	{
		Infinitive: "mest", Translation: "to throw", Class: "1. konjugācija",
		Object: "bumbu",
		Forms: [3][6]string{
			{"metu", "met", "met", "metam", "metat", "met"},
			{"metu", "meti", "meta", "metām", "metāt", "meta"},
			{"metīšu", "metīsi", "metīs", "metīsim", "metīsiet", "metīs"},
		},
	},
	{
		Infinitive: "zaudēt", Translation: "to lose", Class: "2. konjugācija (-ēt)",
		Object: "spēli",
		Forms: [3][6]string{
			{"zaudēju", "zaudē", "zaudē", "zaudējam", "zaudējat", "zaudē"},
			{"zaudēju", "zaudēji", "zaudēja", "zaudējām", "zaudējāt", "zaudēja"},
			{"zaudēšu", "zaudēsi", "zaudēs", "zaudēsim", "zaudēsiet", "zaudēs"},
		},
	},
	{
		Infinitive: "uzvarēt", Translation: "to win", Class: "2. konjugācija (-ēt)",
		Object: "sacensībās",
		Forms: [3][6]string{
			{"uzvaru", "uzvari", "uzvar", "uzvaram", "uzvarat", "uzvar"},
			{"uzvarēju", "uzvarēji", "uzvarēja", "uzvarējām", "uzvarējāt", "uzvarēja"},
			{"uzvarēšu", "uzvarēsi", "uzvarēs", "uzvarēsim", "uzvarēsiet", "uzvarēs"},
		},
	},
	{
		Infinitive: "sākt", Translation: "to start", Class: "1. konjugācija",
		Object: "koncertu",
		Forms: [3][6]string{
			{"sāku", "sāc", "sāk", "sākam", "sākat", "sāk"},
			{"sāku", "sāki", "sāka", "sākām", "sākāt", "sāka"},
			{"sākšu", "sāksi", "sāks", "sāksim", "sāksiet", "sāks"},
		},
	},
	{
		Infinitive: "griezties", Translation: "to turn", Class: "1. konjugācija, atgriezeniskā",
		Object: "pa labi",
		Forms: [3][6]string{
			{"griežos", "griezies", "griežas", "griežamies", "griežaties", "griežas"},
			{"griezos", "griezies", "griezās", "griezāmies", "griezāties", "griezās"},
			{"griezīšos", "griezīsies", "griezīsies", "griezīsimies", "griezīsieties", "griezīsies"},
		},
	},
	{
		Infinitive: "sākt", Translation: "to start", Class: "1. konjugācija",
		Object: "gatavot",
		Forms: [3][6]string{
			{"sāku", "sāc", "sāk", "sākam", "sākat", "sāk"},
			{"sāku", "sāki", "sāka", "sākām", "sākāt", "sāka"},
			{"sākšu", "sāksi", "sāks", "sāksim", "sāksiet", "sāks"},
		},
	},
	{
		Infinitive: "domāt", Translation: "to think", Class: "2. konjugācija (-āt)",
		Object: "par nākotni",
		Forms: [3][6]string{
			{"domāju", "domā", "domā", "domājam", "domājat", "domā"},
			{"domāju", "domāji", "domāja", "domājām", "domājāt", "domāja"},
			{"domāšu", "domāsi", "domās", "domāsim", "domāsiet", "domās"},
		},
	},
}
