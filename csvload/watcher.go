package csvload

const TmpFilesDir = "tmp_files"

// ActivityFileProcessor watches for new input files
// and parses them into activity.
// type ActivityFileProcessor struct {
// 	fileWatcher *fsnotify.Watcher
// 	csvManager  *CSVManager
// 	Done        chan bool
// 	db          *gorm.DB
// 	mux         sync.Mutex
// 	isStarted   bool
// }

// func NewActivityFileProcessor(db *gorm.DB) (*ActivityFileProcessor, error) {
// 	// csvManager, err := csv.NewManager(db)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// fw, err := fsnotify.NewWatcher()

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	processor := &ActivityFileProcessor{
// 		db: db,
// 		// csvManager:  csvManager,
// 		// fileWatcher: fw,
// 		isStarted: false,
// 	}

// 	// db.
// 	// 	Preload("ActivitySource").
// 	// 	Preload("CSVFileConfiguration").
// 	// 	Find(&pipeline.csvSourceMappings)

// 	return processor, nil
// }

// func (processor *ActivityFileProcessor) handleCSVFile(file string) {
// 	// Lookup csv file config
// }

// func (processor *ActivityFileProcessor) startWatchingFiles() {
// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-processor.fileWatcher.Events:
// 				if !ok {
// 					return
// 				}
// 				if event.Op != fsnotify.Create {
// 					continue
// 				}

// 				switch filepath.Ext(event.Name) {
// 				case ".csv":
// 					go processor.handleCSVFile(event.Name)
// 				}
// 			case err, ok := <-processor.fileWatcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				log.Println("error:", err)
// 			}
// 		}
// 	}()

// 	processor.fileWatcher.Add(TmpFilesDir)
// }

// func (processor *ActivityFileProcessor) Start() error {
// 	processor.mux.Lock()
// 	defer processor.mux.Unlock()

// 	if processor.isStarted {
// 		return nil
// 	}

// 	processor.startWatchingFiles()
// 	processor.isStarted = true
// 	return nil
// }

// // func (ap *ActivityPipeline) Stop() {

// // }

// // func (pipeline *ActivityPipeline) process(filename string) {
// // 	dataCh, errCh := pipeline.csvManager.ParseFile(filename)

// // 	handler := func(r csv.Record) {
// // 		// meta, ok := r["_meta"].(map[string]interface{})

// // 		// if !ok {
// // 		// 	log.Println("no meta property on csv record")
// // 		// 	return
// // 		// }

// // 		// fileConfigID, ok := meta["file_config_id"].(uint)

// // 		// if !ok {
// // 		// 	log.Println("no file_config_id on csv record")
// // 		// 	return
// // 		// }

// // 		// // Get activity source mapping
// // 		// var sourceMapping CSVActivitySourceMapping

// // 		// for _, s := range pipeline.csvSourceMappings {
// // 		// 	if s.CSVFileConfigurationID == fileConfigID {
// // 		// 		sourceMapping = s
// // 		// 		break
// // 		// 	}
// // 		// }

// // 		// if sourceMapping.ID == 0 {
// // 		// 	log.Println("no csv activity source mapping found for csv record")
// // 		// }

// // 		// // Create a new activity model from csv record
// // 		// transactionDate, ok := r["Transaction Date"].(time.Time)

// // 		// if !ok {
// // 		// 	log.Fatal("cannot transform transaction date")
// // 		// }

// // 		// postedDate, ok := r["Posted Date"].(time.Time)

// // 		// if !ok {
// // 		// 	log.Fatal("cannot transform posted date")
// // 		// }

// // 		// desc, ok := r["Description"]

// // 		// if !ok {
// // 		// 	log.Fatal("cannot transform posted date")
// // 		// }

// // 		// amount, ok := r["Amount"]
// // 	}

// // 	for {
// // 		// blocks until there's data available on ch1 or ch2
// // 		select {
// // 		case data, ok := <-dataCh:
// // 			if !ok {
// // 				return
// // 			}
// // 			handler(data)
// // 		case err := <-errCh:
// // 			fmt.Println(err)
// // 		}
// // 	}
// // }
