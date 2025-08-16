package store

const (
	UZ string = "uz"
	RU string = "ru"
	EN string = "en"
)

const (
	LanguageUzLabel    string = "language_uz"
	LanguageRuLabel    string = "language_ru"
	LanguageEnLabel    string = "language_en"
	AdvertisementLabel string = "advertisement"
	AdviceLabel        string = "advice"
	FaqLabel           string = "faq"
	SuggestionLabel    string = "suggestion"
	YoutubeLabel       string = "youtube"
	TelegramLabel      string = "telegram"
	InstagramLabel     string = "instagram"
	FacebookLabel      string = "facebook"
	ReturnLabel        string = "return"
)

var ButtonLabels = map[string]map[string]string{
	LanguageUzLabel:    {UZ: "🇺🇿 O‘zbekcha", RU: "🇺🇿 Узбекский", EN: "🇺🇿 Uzbek"},
	LanguageRuLabel:    {UZ: "🇷🇺 Ruscha", RU: "🇷🇺 Русский", EN: "🇷🇺 Russian"},
	LanguageEnLabel:    {UZ: "🇬🇧 Inglizcha", RU: "🇬🇧 Английский", EN: "🇬🇧 English"},
	AdvertisementLabel: {UZ: "🎯 REKLAMA", RU: "🎯 РЕКЛАМА", EN: "🎯 ADVERTISEMENT"},
	AdviceLabel:        {UZ: "💡 MASLAHAT", RU: "💡 СОВЕТ", EN: "💡 ADVICE"},
	FaqLabel:           {UZ: "🤔 KO'P BERILADIGAN SAVOLLAR", RU: "🤔 ЧАСТЫЕ ВОПРОСЫ", EN: "🤔 FAQ"},
	SuggestionLabel:    {UZ: "📬 TAKLIF QILISH", RU: "📬 ПРЕДЛОЖЕНИЕ", EN: "📬 SUGGESTION"},
	YoutubeLabel:       {UZ: "▶️ YOUTUBE", RU: "▶️ YOUTUBE", EN: "▶️ YOUTUBE"},
	TelegramLabel:      {UZ: "✈️ TELEGRAM", RU: "✈️ TELEGRAM", EN: "✈️ TELEGRAM"},
	InstagramLabel:     {UZ: "📸 INSTAGRAM", RU: "📸 INSTAGRAM", EN: "📸 INSTAGRAM"},
	FacebookLabel:      {UZ: "👥 FACEBOOK", RU: "👥 FACEBOOK", EN: "👥 FACEBOOK"},
	ReturnLabel:        {UZ: "⬅️ ORQAGA", RU: "⬅️ НАЗАД", EN: "⬅️ BACK"},
}
