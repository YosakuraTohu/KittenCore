package track

type (
	// 状态
	status byte

	// 状态错误
	statusErr struct {
		url  string // 链接
		stat status // 状态
	}

	// 不支持的平台
	notSupportedErr struct {
		p platform // 平台
	}

	// 没有找到小说
	notFoundErr struct {
		key keyword // 搜索关键词
	}
)

const (
	bookUnreachable        status = iota // 无法访问小说
	noChapterURL                         // 没有章节链接
	onlyAChapter                         // 只有一个章节
	bookStatusException                  // 小说状态异常
	timeException                        // 更新时间异常
	chapterUnreachable                   // 无法访问章节
	chapterURLException                  // 章节链接异常
	chapterStatusException               // 章节状态异常
	vipChapterException                  // 付费状态异常
)

// Error 实现 error
func (e *statusErr) Error() string {
	switch e.stat {
	case bookUnreachable:
		return `链接 ` + e.url + ` 没有小说喵！`
	case noChapterURL:
		return `小说 ` + e.url + ` 没有章节链接喵！`
	case onlyAChapter:
		return `小说 ` + e.url + ` 只有一个章节喵！`
	case bookStatusException:
		return `小说 ` + e.url + ` 状态异常喵！`
	case timeException:
		return `小说 ` + e.url + ` 上次更新时间异常喵！`
	case chapterUnreachable:
		return `链接 ` + e.url + ` 没有章节喵！`
	case chapterURLException:
		return e.url + ` 不是正常的章节链接喵！`
	case chapterStatusException:
		return `章节 ` + e.url + ` 状态异常喵！`
	case vipChapterException:
		return `章节 ` + e.url + ` 付费状态异常喵！`
	default:
		return `状态错误`
	}
}

// *statusErr 的构造函数，状态错误
func errStatus(url string, stat status) *statusErr {
	return &statusErr{
		url:  url,
		stat: stat,
	}
}

// *notSupportedErr 的构造函数，不支持的平台
func notSupported(p platform) *notSupportedErr {
	return &notSupportedErr{
		p: p,
	}
}

// Error 实现 error
func (e *notSupportedErr) Error() string {
	return string(e.p) + ` 不是受支持的小说平台喵！`
}

// *notFoundErr 的构造函数，没有找到小说
func notFound(key keyword) *notFoundErr {
	return &notFoundErr{
		key: key,
	}
}

// Error 实现 error
func (e *notFoundErr) Error() string {
	return `没有找到` + string(e.key) + `关键词的小说喵！`
}
