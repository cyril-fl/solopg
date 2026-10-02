package initI19n

import (
	"fmt"
	"path/filepath"
	"slices"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/process"
	"solopg/app/shared/services/yaml"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	y "gopkg.in/yaml.v3"
)

type initializer struct {
	process.Process
	cache

	cfg    i19n.Config
	locale string
}

type cache struct {
	defaultLocale *i19n.Locale

	bundle     *i18n.Bundle
	local      *i18n.Localizer
	localerror *i18n.Localizer

	files []string
}

func Process(cfg i19n.Config, locale string) *initializer {
	return &initializer{
		cfg:    cfg,
		locale: locale,
	}
}

func (i *initializer) Run() {
	// Set env
	i.reset()
	i.assertConfig()
	i.setDefaultLocale()

	// Bundle
	i.initBundle()
	i.loadLocaleFiles()
	i.registerLocaleFiles()

	// Localizer
	i.initLocalizer()
	i.initErrorLocalizer()

	// Set
	i.seti19n()
}

func (i *initializer) GetResult() {}

// Methods
func (i *initializer) assertConfig() {
	if i.HasErr() {
		return
	}

	if err := i.cfg.Validate(); err != nil {
		i.SetErr(err)
	}
}

func (i *initializer) setDefaultLocale() {
	if i.HasErr() {
		return
	}

	defaultLocale, err := i.cfg.GetDefaultLocale(i.cfg.Default)
	if err != nil {
		i.SetErr(err)
		return
	}

	i.defaultLocale = defaultLocale
}

func (i *initializer) initBundle() {
	if i.HasErr() {
		return
	}

	tag, err := i.defaultLocale.ParseTag()
	if err != nil {
		i.SetErr(err)
		return
	}

	bundle := i18n.NewBundle(tag)
	bundle.RegisterUnmarshalFunc(string(i.cfg.Format), y.Unmarshal)

	i.bundle = bundle
}

func (i *initializer) loadLocaleFiles() {
	if i.HasErr() {
		return
	}

	files, err := loadLocaleFile(i.cfg)
	if err != nil {

		i.SetErr(err)
		return
	}

	i.files = files
}

func (i *initializer) registerLocaleFiles() {
	if i.HasErr() {
		return
	}

	for _, path := range i.files {
		if _, err := i.bundle.LoadMessageFile(path); err != nil {
			i.SetErr(fmt.Errorf("load message file '%s': %w", path, err))
		}
	}
}

func (i *initializer) initLocalizer() {
	if i.HasErr() {
		return
	}

	localizer, err := newLocalizer(i.cfg, i.bundle, i.locale, i.defaultLocale)
	if err != nil {
		i.SetErr(err)
		return
	}

	i.local = localizer
}

func (i *initializer) initErrorLocalizer() {
	if i.HasErr() {
		return
	}

	localizer, err := newLocalizer(i.cfg, i.bundle, "en", i.defaultLocale)
	if err != nil {
		i.SetErr(err)
		return
	}

	i.localerror = localizer
}

func (i *initializer) seti19n() {
	if i.HasErr() {
		return
	}

	t := i19n.New()
	t.SetBundle(i.bundle)
	t.SetLocale(i.local)
	t.SetLocalerror(i.localerror)

	i19n.SetCache(t)
}

func (i *initializer) reset() {
	i.cache = cache{}

	i.SetErr(nil)
	i19n.SetCache(i19n.New())
}

// Helpers
func loadLocaleFile(cfg i19n.Config) ([]string, error) {
	files, err := yaml.GetFilesFromSource(cfg.Dir, true)
	if err != nil {
		return nil, err
	}

	assertedFiles := make([]string, 0)
	for _, path := range files {
		extensions := i19n.FormatValidExtensions[cfg.Format]

		if !(slices.Contains(extensions, filepath.Ext(path))) {
			continue
		}

		name := filepath.Base(path)
		name = strings.TrimSuffix(name, filepath.Ext(name))
		if _, err := cfg.GetLocaleByISO(name); err != nil {
			continue
		}

		assertedFiles = append(assertedFiles, path)
	}

	return assertedFiles, nil
}

/*
NOTE newLocalizer pourrais très biens passer en methode mais je treouve que c'est mélanger les responsabiltés.
Ainsi un localierpeu etre initialisé sans passer par l'initilizer.
*/
func newLocalizer(cfg i19n.Config, bundle *i18n.Bundle, lang string, defaultLocale *i19n.Locale) (*i18n.Localizer, error) {
	locale, err := getLocale(cfg, lang, defaultLocale)
	if err != nil {
		return nil, err
	}

	tag, err := locale.ParseTag()
	if err != nil {
		return nil, err
	}

	defaultTag, err := defaultLocale.ParseTag()
	if err != nil {
		return nil, err
	}

	return i18n.NewLocalizer(bundle, tag.String(), defaultTag.String()), nil
}

func getLocale(cfg i19n.Config, lang string, defaultLocale *i19n.Locale) (*i19n.Locale, error) {
	if lang == "" {
		return defaultLocale, nil
	}

	locale, err := cfg.GetLocaleByCode(lang)
	if err != nil {
		return nil, fmt.Errorf("locale '%s' not found: %w", lang, err)
	}

	return locale, nil
}
