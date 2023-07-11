package initializers
// Initialize logger
import(
	log "github.com/sirupsen/logrus"
	"os"
)
func InitLogger() {
	// Log as JSON instead of the default ASCII formatter
	log.SetFormatter(&log.JSONFormatter{})

	// Set log level
	log.SetLevel(log.DebugLevel)
	log.SetLevel(log.ErrorLevel)
	log.SetLevel(log.ErrorLevel)
	log.SetLevel(log.WarnLevel)
	log.SetLevel(log.TraceLevel)

	// Create or append to log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}