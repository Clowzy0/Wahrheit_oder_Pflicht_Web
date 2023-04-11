package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"errors"
	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type phrase struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Level   string `json:"level"`
	Content string `json:"content"`
	TFT     int    `json:"tft"` //TFT = Time For Task (seconds)
	RAP     bool   `json:"rap"`
	TTS     int    `json:"tts"`
}

var phrases = []phrase{
	{ID: "0", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was war der komischste Traum den du jemals hattest?"},
	{ID: "1", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "hast du ein Crush?"},
	{ID: "2", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was war die größte Lüge die du jemals erzählt hast?"},
	{ID: "3", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was ist ein komisches Gericht dass du liebst?"},
	{ID: "4", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was ist das schlimmste dass du je getan hast?"},
	{ID: "5", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "welche deiner Beziehungen hatte den größten Altersunterschied und wie groß war dieser?"},
	{ID: "6", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "willst du jemals heiraten?"},
	{ID: "7", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "willst du jemals ein oder mehr Kinder haben?"},
	{ID: "8", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was war die größte Geldverschwendung in deinem Leben?"},
	{ID: "9", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "hast du ein verstecktes Talent?"},
	{ID: "10", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "gibt es eine Person bei der du so getan hast als ob du sie magst wenn ja wer und warum?"},
	{ID: "11", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "hast du jemals gegen das Gesetzt verstoßen wenn ja was hast du gemacht?"},
	{ID: "12", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "wer im Raum ist deiner Meinung nach am schlechtesten gekleidet?"},
	{ID: "13", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was ist das seltsamste Gerücht über dich selbst dass du je gehört hast?"},
	{ID: "14", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "hast du jemals gelogen um aus einem schlechten Date zu entkommen wenn ja was war die Lüge?"},
	{ID: "15", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "hast du jemals jemandem gesagt dass du nicht Zuhause bist weil du dich nicht mit der Person treffen wolltest?"},
	{ID: "16", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was war der ekelhafteste Ort an dem du je warst?"},
	{ID: "17", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was ist das ekelhafteste dass du je getrunken hast?"},
	{ID: "18", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "welche Sache in deinem Leben bereust du am meisten?"},
	{ID: "19", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "wann hast du das letze mal geweint?"},
	{ID: "20", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "wenn du 3 Wünsche hättest welche wären diese?"},
	{ID: "21", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "wenn du eine Sache ohne Folgen tun könntest was würdest du machen?"},
	{ID: "22", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was ist deine größte Angst?"},
	{ID: "23", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "was ist deine größte Unsicherheit?"},
	{ID: "24", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "welche Lüge hast du zuletzt erzählt?"},
	{ID: "25", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "wann hast du zuletzt geduscht?"},
	{ID: "26", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "pinkeltst du in der Dusche?"},
	{ID: "27", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "hörst du beim Duschen Musik?"},
	{ID: "28", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "SFW", Content: "muss dein Partner gut aussehen?"},
	//***"2*"**RAP: false, *********************************//
	{ID: "29", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du schonmal Drogen genommen wenn ja welche?"},
	{ID: "30", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "wann hast du zum letzten mal etwas das NSFW war geschaut/gelsen/konsumiert und was?"},
	{ID: "31", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hattest du schonmal Sex wenn ja wie war es?"},
	{ID: "32", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du ein Fetisch wenn ja welchen?"},
	{ID: "33", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du jemals jemanden betrogen?"},
	{ID: "34", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "was ist für dich der größte Turn-off?"},
	{ID: "35", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du jemals ein Orgasmus vorgetäuscht?"},
	{ID: "36", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "wurdest du jemals beim Sex erwischt?"},
	{ID: "37", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "ist Penisgröße für dich wichtig?"},
	{ID: "38", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "ist Körpergröße für dich wichtig?"},
	{ID: "39", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "ist Körbchen-/Brustgröße für dich wichtig?"},
	{ID: "40", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "sind Muskeln für dich wichtig?"},
	{ID: "41", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "wenn du nur eine Sexposition für den Rest deines Lebens verwenden könntest welche wäre diese?"},
	{ID: "42", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du je ein Sextape gefilmt?"},
	{ID: "43", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "was ist der größte Turn-on für dich?"},
	{ID: "44", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "was ist deine verrückteste ONS Erfahrung?"},
	{ID: "45", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "mit wem hattest du deinen schlechtesten Kuss?"},
	{ID: "46", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du jemals deine Eltern beim Sex erwischt?"},
	{ID: "47", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "wann warst du am meisten betrunken?"},
	{ID: "48", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "was wär das komischste dass jemand in deinem Suchverlauf sehen würde?"},
	{ID: "49", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du jemals ein Blowjob gegeben wenn ja wie war es?"},
	{ID: "50", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du schonmal Sperma probiert wenn ja wie war es?"},
	{ID: "51", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du jemal ein Boob-/Titjob gegeben wenn ja wie war es?"},
	{ID: "52", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du jemals Analsex probiert wenn ja wie war es?"},
	{ID: "53", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du schonmal deepthroating probiert wenn ja wie wars?"},
	{ID: "54", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "wenn du mit einer Person aus dem Raum schlafen müssstest wer wäre es?"},
	{ID: "55", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "wenn du mit einer Person schlafen müssstest wer wäre diese?"},
	{ID: "56", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du Sextoys?"},
	{ID: "57", RAP: false, TFT: 0, TTS: 0, Type: "Truth", Level: "NSFW", Content: "hast du schonmal mastrubiert wenn ja hat es ich gut angefühlt?"},
	//***"5*"**RAP: false, *********************************//
	{ID: "58", RAP: false, TFT: 120, TTS: 40, Type: "Dare", Level: "SFW", Content: "hol etwas aus dem Badezimmer und versuche 2 Minuten lang es den anderen zu verkaufen."},
	{ID: "59", RAP: true, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "lass eine person aus dem Raum ein Tatto mit einem Marker irgendwo auf deinen Körper malen."},
	{ID: "60", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "die anderen müssen ein Getränk vorbereiten welches du dann trinken musst."},
	{ID: "61", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "lasse 2 Eiswürfel in deinem Mund bis sie schmelzen."},
	{ID: "62", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "schrei das erste Wort dass dir in den Sinn kommt."},
	{ID: "63", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "versuch etwas zu essen dass die anderen dir geben ohne deine Hände zu benutzen."},
	{ID: "64", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "SCHREI!!!"},
	{ID: "65", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "sage je 2 Ehrliche Dinge über jede Person die im Raum ist."},
	{ID: "66", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "werf ein Eiswürfel in deine Hose."},
	{ID: "67", RAP: false, TFT: 7, TTS: 10, Type: "Dare", Level: "SFW", Content: "versuche 7 sekunden lang ein Breakdance vorzuführen."},
	{ID: "68", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "geb jedem im Raum eine Persönliche Beleidigung."},
	{ID: "69", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "rufe eine zufällige Telefonnummer an und versuche so lange wie möglich mit der Person zu reden."},
	{ID: "70", RAP: false, TFT: 30, TTS: 3, Type: "Dare", Level: "SFW", Content: "streite dich für 30 Sekunden mit einer wand."},
	{ID: "71", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "trink ein Shot essig."},
	{ID: "72", RAP: true, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "blockiere deine Augen und lasse dir von einer person aus dem Raum etwas in den Mund legen dass du dann essen musst."},
	{ID: "73", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "beschreibe eine ekelhafte Gewohnheit von dir."},
	{ID: "74", RAP: true, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "du musst dich für die nächsten 5 minuten neben eine person aus dem Raum setzen und ihn/sie cuddlen z.B. Arm umklammern, Kopf auf die Schulter legen etc."},
	{ID: "75", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "lass die anderen im Raum eine Minute lang durch dein Handy schauen."},
	{ID: "76", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "zeige die letzen 3 Tage deines Suchverlaufs."},
	{ID: "77", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "schreibe deinem/deiner Ex dass du ihn/sie liebst."},
	{ID: "78", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "umarme eine Person deiner Wahl aus dem Raum egal wie du willst."},
	{ID: "79", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "tausche ein Kleidungsstück mit der Person links von dir."},
	{ID: "80", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "SFW", Content: "tausche ein Kleidungsstück mit der Person rechts von dir."},
	//***"8*"**RAP: false, *********************************//
	{ID: "81", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "küsse eine andere Person aus dem Raum auf die Wange."},
	{ID: "82", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "küsse eine andere Person aus dem Raum auf den Mund."},
	{ID: "83", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "küsse eine andere Person aus dem Raum wohin du willst."},
	{ID: "84", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "führe eine Sexposition deiner Wahl vor."},
	{ID: "85", RAP: false, TFT: 7, TTS: 0, Type: "Dare", Level: "NSFW", Content: "mache mit der Person rechts von dir Augenkontakt und moan für 7 Sekunden."},
	{ID: "86", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "führe einen Unterwürfigen Tanz vor der Gruppe vor."},
	{ID: "87", RAP: false, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "nehme die nähste Sache in deiner Reichweite und demonstriere daran wie man ein Kondom aufzieht."},
	{ID: "88", RAP: true, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "küsse eine person aus dem Raum auf die Wange."},
	{ID: "89", RAP: true, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "küsse eine person aus dem Raum den Mund."},
	{ID: "90", RAP: true, TFT: 0, TTS: 0, Type: "Dare", Level: "NSFW", Content: "küsse eine person aus dem Raum wohin du willst."},
}

func get_all_phrases(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, phrases)
}

func phrase_by_ID(c *gin.Context) {
	id := c.Param("id")
	phrase, err := get_phrase_by_ID(id)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, phrase)
}

func get_phrase_by_ID(id string) (*phrase, error) {
	for i, phrase := range phrases {
		if phrase.ID == id {
			return &phrases[i], nil
		}
	}

	return nil, errors.New("phrase not found")
}

func get_SFW_truth(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 28

	phrase_id := (rand.Intn(max-min) + min)
	c.IndentedJSON(http.StatusOK, phrases[phrase_id])
}

func get_ANY_truth(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 57

	phrase_id := (rand.Intn(max-min) + min)
	c.IndentedJSON(http.StatusOK, phrases[phrase_id])
}

func get_SFW_dare(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	min := 58
	max := 80

	phrase_id := (rand.Intn(max-min) + min)
	c.IndentedJSON(http.StatusOK, phrases[phrase_id])
}

func get_ANY_dare(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	min := 58
	max := 90

	phrase_id := (rand.Intn(max-min) + min)
	c.IndentedJSON(http.StatusOK, phrases[phrase_id])
}

func save_phrase(c *gin.Context) {
	var newPhrase phrase

	if err := c.BindJSON(&newPhrase); err != nil {
		return
	}

	phrases = append(phrases, newPhrase)
	c.IndentedJSON(http.StatusCreated, newPhrase)
}

func save_all_phrases(c *gin.Context) {
	file, _ := json.MarshalIndent(phrases, "", " ")

	_ = ioutil.WriteFile("phrases.json", file, 0644)
}

func main() {
	fmt.Println("Starting Gin in ReleaseMode...")
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("Ready!")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/dump_phrases", get_all_phrases)
	router.GET("/get_phrase/:id", phrase_by_ID)
	router.GET("/get_ANY_dare", get_ANY_dare)
	router.GET("/get_SFW_dare", get_SFW_dare)
	router.GET("/get_ANY_truth", get_ANY_truth)
	router.GET("/get_SFW_truth", get_SFW_truth)
	router.GET("/save_all_phrases", save_all_phrases)
	router.Run("localhost:1234")
}
