package donation

//var InstanceChat *Chat
//
//func Start() {
//	//чат крутится как отдельная сущность всегда
//	InstanceChat = NewChat(config.Get().ChatConfig.MaxUsers)
//	go panicWorker.PanicWorker(InstanceChat.Start)
//}
//
//type Notification struct {
//	mu           *sync.Mutex
//	UserChan     chan *User
//	Users        map[string]*User
//	MaxUsers     int
//	register     chan *User
//	unregister   chan *User
//	ticker       *time.Ticker
//	messagesChan chan UserMessage
//	messages     []Message
//}
//
//func NewChat(maxUsers int) *Chat {
//	return &Chat{
//		mu:           &sync.Mutex{},
//		UserChan:     make(chan *User, 10),
//		Users:        make(map[string]*User),
//		MaxUsers:     maxUsers,
//		register:     make(chan *User, 10),
//		unregister:   make(chan *User, 10),
//		ticker:       time.NewTicker(1 * time.Second),
//		messagesChan: make(chan UserMessage, 10),
//		messages:     make([]Message, 0, 100),
//	}
//}
