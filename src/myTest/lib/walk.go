package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

func CountFiles() int {
	count := 0
	imagesFiles := 0
	ext := []string{".bin"}

	//rootPath := "/home/kasper/Downloads/documents/"
	//rootPath := "/home/kasper/Downloads/"
	rootPath := "/home/kasper/"

	filepath.WalkDir(rootPath, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !file.IsDir() {
			file.Type()
			//if !file.IsDir() {

			//fmt.Println(path)
			//info := file.Type()
			//if err != nil {
			//	return nil
			//}

			//fmt.Printf("variable info.Name() = %v is of type %T \n", info.Type(), info.Type())

			//".jpg", ".jpeg", ".png", ".gif", ".webp", ".tiff", ".tif", ".psd", ".raw", ".arw", ".cr2", ".nrw", ".k25", ".bmp", ".dib", ".heif", ".heic", ".ind", ".indd", ".indt", ".svg", ".svgz", ".ai", ".eps", ".pdf"

			images := []string{".jpg", ".bmp", ".gif", ".png", ".ttf", ".PNG"}

			//test := []string{".bin", ".swp", ".js", ".css", ".map", ".svg", ".pma", ".old", ".log", ".db",
			//	".sh", ".md", ".yml", ".py", ".rst", ".txt", ".pem", ".j2", ".pyc", ".cfg",
			//	".iso", ".p12", ".crt", ".key", ".hcl", ".gz", ".c", ".xz", ".sql", ".tf",
			//	".xml", ".atd", ".com", ".ini", ".pak", ".dat", ".so", ".1", ".mjs", ".cjs",
			//	".ts", ".gyp", ".d", ".mk", ".env", ".js~", ".sh~", ".ts~", ".fnt", ".bnf",
			//	".ejs", ".ico", ".jsc", ".mts", ".jst", ".def", ".iml", ".mkd", ".md~",
			//	".njs", ".in", ".BSD", ".bak", ".hbs", ".l", ".y", ".nix", ".mli", ".ps1",
			//	".swf", ".un~", ".MIT", ".exe", ".ml", ".tsx", ".re", ".pid", ".len", ".at",
			//	".s", ".ver", ".csv", ".939", ".zip", ".ij", ".bpe", ".lck", ".174", ".622",
			//	".973", ".8", ".29", ".10", ".8_i", ".2", ".2_i", ".tab", ".0", ".125", ".150",
			//	".154", ".159", ".248", ".33", ".375", ".384", ".43", ".494", ".531", ".616",
			//	".625", ".629", ".647", ".817", ".959", ".27", ".ic", ".tmp", ".106", ".12",
			//	".13", ".184", ".201", ".203", ".209", ".215", ".227", ".262", ".276", ".285",
			//	".289", ".297", ".324", ".338", ".339", ".345", ".409", ".412", ".447", ".490",
			//	".50", ".500", ".540", ".556", ".563", ".584", ".600", ".61", ".610", ".614",
			//	".651", ".666", ".697", ".702", ".710", ".711", ".715", ".747", ".751", ".755",
			//	".758", ".774", ".787", ".808", ".815", ".830", ".845", ".853", ".865", ".866",
			//	".87", ".874", ".882", ".885", ".887", ".888", ".919", ".928", ".947", ".966",
			//	".975", ".976", ".995", ".996", ".xd", ".tr", ".107", ".108", ".115", ".14",
			//	".146", ".168", ".17", ".179", ".186", ".189", ".193", ".202", ".263", ".264",
			//	".306", ".31", ".328", ".348", ".354", ".368", ".373", ".387", ".389", ".392",
			//	".404", ".407", ".415", ".416", ".426", ".430", ".434", ".462", ".471", ".480",
			//	".483", ".5", ".503", ".507", ".524", ".530", ".537", ".544", ".547", ".557",
			//	".562", ".578", ".581", ".583", ".617", ".634", ".639", ".65", ".67", ".68",
			//	".680", ".688", ".693", ".725", ".73", ".730", ".767", ".769", ".771", ".778",
			//	".779", ".785", ".790", ".797", ".804", ".806", ".812", ".814", ".819", ".820",
			//	".846", ".848", ".854", ".870", ".890", ".894", ".903", ".911", ".92", ".937",
			//	".954", ".96", ".968", ".971", ".98", ".212", ".788", ".701", ".881", ".112",
			//	".xb", ".eet", ".TAG", ".h", ".lz4", ".tgz", ".cmd", ".cc", ".php", ".el",
			//	".jsx", ".bar", ".wmf", ".lcl", ".dmp", ".go", ".ggr", ".jar", ".pom", ".pb",
			//	".wat", ".mp3", ".mp4", ".otf", ".eot", ".ldb", ".fbs", ".MF", ".mem", ".hyb",
			//	".bau", ".xba", ".xlb", ".xlc", ".fmt", ".dbf", ".dbt", ".odb", ".sdv", ".thm",
			//	".xcu", ".tdb", ".crl", ".kbx", ".gpg", ".sym", ".jfc", ".ja", ".pyi", ".dll",
			//	".pdb", ".dir", ".49", ".lst", ".dtd", ".6", ".3", ".11", ".15", ".pdf", ".238",
			//	".ijx", ".zsh", ".240", ".bat", ".102", ".egg", ".241", ".pp", ".rb", ".whl",
			//	".tex", ".sty", ".odt", ".htc", ".pxd", ".pyx", ".pyd", ".hpp", ".cpp", ".pyz",
			//	".47", ".ods", ".idx", ".deb", ".vue", ".xcf", ".mdx", ".kwl", ".rc", ".qml",
			//	".zo", ".7", ".cs", ".cts", ".jsa", ".pub", ".msf", ".grd", ".ppt", ".eps",
			//	".MD", ".cnf", ".crx", ".mod", ".sum", ".4", ".di", ".uml", ".eml", ".srl",
			//	".git", ".jsp", ".inc"}

			if slices.Contains(images, filepath.Ext(path)) {
				imagesFiles++
			} else {
				if !slices.Contains(ext, strings.ToLower(filepath.Ext(path))) && len(filepath.Ext(path)) < 5 {
					ext = append(ext, filepath.Ext(path))
				}
			}
			count++
		}

		return nil
	})

	fmt.Printf("variable images = %v is of type %T \n", imagesFiles, imagesFiles)
	fmt.Printf("variable ext = %v is of type %T \n", ext, ext)

	return count
}
