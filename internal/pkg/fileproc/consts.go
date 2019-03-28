package fileproc

const MaxUploadSize = 2 * 1024 * 1024 // 2 mb
const UploadPath = "/var/www/media/factorial"

type StaticServerConfig struct {
	MaxUploadSizeMB int64  `json:"max_upload_size_mb"`
	UploadPath      string `json:"upload_path"`
	MaxUploadSize   int64
}

//var StaticConfig = StaticServerConfig{}

//func init() {
//	err := config_reader.ReadConfigFile("static_server_config.json", &StaticConfig)
//	if err != nil {
//		log.Fatal(errors.Wrap(err, "error while reading Cookie config"))
//	}
//
//	StaticConfig.MaxUploadSize = StaticConfig.MaxUploadSizeMB * 1024 * 1024
//
//	fmt.Println(StaticConfig)
//}
