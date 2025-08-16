package store

type Store struct{}

func New() Store {
	return Store{}
}

const (
	UZ string = "uz"
	RU string = "ru"
	EN string = "en"
)

const (
	StateStartCommand = iota
	StateLanguage
	StateMenu
	StateAdvertisement
	StateAdvice
	StateSuggestion
)

const (
	LabelLanguageUz string = "language_uz"
	LabelLanguageRu string = "language_ru"
	LabelLanguageEn string = "language_en"

	LabelAdvertisement string = "advertisement"
	LabelAdvice        string = "advice"
	LabelFaq           string = "faq"
	LabelSuggestion    string = "suggestion"
	LabelYoutube       string = "youtube"
	LabelTelegram      string = "telegram"
	LabelInstagram     string = "instagram"
	LabelFacebook      string = "facebook"

	LabelPrice      string = "price"
	LabelStatistics string = "statistics"

	LabelFree string = "free"
	LabelPaid string = "paid"

	LabelPhone string = "phone"

	LabelReturn string = "return"
)

var ButtonLabels = map[string]map[string]string{
	LabelLanguageUz: {UZ: "🇺🇿 O‘zbekcha", RU: "🇺🇿 Узбекский", EN: "🇺🇿 Uzbek"},
	LabelLanguageRu: {UZ: "🇷🇺 Ruscha", RU: "🇷🇺 Русский", EN: "🇷🇺 Russian"},
	LabelLanguageEn: {UZ: "🇬🇧 Inglizcha", RU: "🇬🇧 Английский", EN: "🇬🇧 English"},

	LabelAdvertisement: {UZ: "🎯 REKLAMA", RU: "🎯 РЕКЛАМА", EN: "🎯 ADVERTISEMENT"},
	LabelAdvice:        {UZ: "💡 MASLAHAT", RU: "💡 СОВЕТ", EN: "💡 ADVICE"},
	LabelFaq:           {UZ: "🤔 KO'P BERILADIGAN SAVOLLAR", RU: "🤔 ЧАСТЫЕ ВОПРОСЫ", EN: "🤔 FAQ"},
	LabelSuggestion:    {UZ: "📬 TAKLIF QILISH", RU: "📬 ПРЕДЛОЖЕНИЕ", EN: "📬 SUGGESTION"},
	LabelYoutube:       {UZ: "▶️ YOUTUBE", RU: "▶️ YOUTUBE", EN: "▶️ YOUTUBE"},
	LabelTelegram:      {UZ: "📨 TELEGRAM", RU: "📨 TELEGRAM", EN: "📨 TELEGRAM"},
	LabelInstagram:     {UZ: "📸 INSTAGRAM", RU: "📸 INSTAGRAM", EN: "📸 INSTAGRAM"},
	LabelFacebook:      {UZ: "👥 FACEBOOK", RU: "👥 FACEBOOK", EN: "👥 FACEBOOK"},

	LabelPrice:      {UZ: "🏷️ NARX", RU: "🏷️ ЦЕНА", EN: "🏷️ PRICE"},
	LabelStatistics: {UZ: "📊 STATISTIKA", RU: "📊 СТАТИСТИКА", EN: "📊 STATISTICS"},

	LabelFree: {UZ: "🎁 BEPUL", RU: "🎁 БЕСПЛАТНО", EN: "🎁 FREE"},
	LabelPaid: {UZ: "💳 PULLIK", RU: "💳 ПЛАТНО", EN: "💳 PAID"},

	LabelPhone: {UZ: "📞 TELEFON", RU: "📞 ТЕЛЕФОН", EN: "📞 PHONE"},

	LabelReturn: {UZ: "⬅️ ORQAGA", RU: "⬅️ НАЗАД", EN: "⬅️ BACK"},
}
