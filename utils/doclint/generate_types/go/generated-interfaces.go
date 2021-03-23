package playwright
/// <summary>
	/// <para>
	/// A Browser is created via <see cref="BrowserType.Launch"/>. An example of using a
	/// <see cref="Browser"/> to create a <see cref="Page"/>:
	/// </para>
	/// </summary>
	type Browser interface {
	EventEmitter
		/// <summary>
		/// <para>
		/// Emitted when Browser gets disconnected from the browser application. This might
		/// happen because of one of the following:
		/// </para>
		/// <list type="bullet">
		/// <item><description>Browser application is closed or crashed.</description></item>
		/// <item><description>The <see cref="Browser.Close"/> method was called.</description></item>
		/// </list>
		/// </summary>
		event EventHandler Disconnected;
		/// <summary>
		/// <para>
		/// In case this browser is obtained using <see cref="BrowserType.Launch"/>, closes
		/// the browser and all of its pages (if any were opened).
		/// </para>
		/// <para>
		/// In case this browser is connected to, clears all created contexts belonging to this
		/// browser and disconnects from the browser server.
		/// </para>
		/// <para>
		/// The <see cref="Browser"/> object itself is considered to be disposed and cannot
		/// be used anymore.
		/// </para>
		/// </summary>
		Close() void
		/// <summary>
		/// <para>
		/// Returns an array of all open browser contexts. In a newly created browser, this
		/// will return zero browser contexts.
		/// </para>
		/// </summary>
		IReadOnlyCollection<BrowserContext> Contexts { get; }
		/// <summary><para>Indicates that the browser is connected.</para></summary>
		*bool IsConnected { get; }
		/// <summary><para>Creates a new browser context. It won't share cookies/cache with other browser contexts.</para></summary>
		/// <param name="acceptDownloads">
		/// Whether to automatically download all the attachments. Defaults to <c>false</c>
		/// where all the downloads are canceled.
		/// </param>
		/// <param name="bypassCSP">Toggles bypassing page's Content-Security-Policy.</param>
		/// <param name="colorScheme">
		/// Emulates <c>'prefers-colors-scheme'</c> media feature, supported values are <c>'light'</c>,
		/// <c>'dark'</c>, <c>'no-preference'</c>. See <see cref="Page.EmulateMedia"/> for more
		/// details. Defaults to <c>'light'</c>.
		/// </param>
		/// <param name="deviceScaleFactor">Specify device scale factor (can be thought of as dpr). Defaults to <c>1</c>.</param>
		/// <param name="extraHTTPHeaders">
		/// An object containing additional HTTP headers to be sent with every request. All
		/// header values must be strings.
		/// </param>
		/// <param name="geolocation">
		/// </param>
		/// <param name="hasTouch">Specifies if viewport supports touch events. Defaults to false.</param>
		/// <param name="httpCredentials">
		/// Credentials for <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication">HTTP
		/// authentication</a>.
		/// </param>
		/// <param name="ignoreHTTPSErrors">Whether to ignore HTTPS errors during navigation. Defaults to <c>false</c>.</param>
		/// <param name="isMobile">
		/// Whether the <c>meta viewport</c> tag is taken into account and touch events are
		/// enabled. Defaults to <c>false</c>. Not supported in Firefox.
		/// </param>
		/// <param name="javaScriptEnabled">Whether or not to enable JavaScript in the context. Defaults to <c>true</c>.</param>
		/// <param name="locale">
		/// Specify user locale, for example <c>en-GB</c>, <c>de-DE</c>, etc. Locale will affect
		/// <c>navigator.language</c> value, <c>Accept-Language</c> request header value as
		/// well as number and date formatting rules.
		/// </param>
		/// <param name="offline">Whether to emulate network being offline. Defaults to <c>false</c>.</param>
		/// <param name="permissions">
		/// A list of permissions to grant to all pages in this context. See <see cref="BrowserContext.GrantPermissions"/>
		/// for more details.
		/// </param>
		/// <param name="proxy">
		/// Network proxy settings to use with this context. Note that browser needs to be launched
		/// with the global proxy for this option to work. If all contexts override the proxy,
		/// global proxy will be never used and can be any string, for example <c>launch({ proxy:
		/// { server: 'per-context' } })</c>.
		/// </param>
		/// <param name="recordVideo">
		/// Enables video recording for all pages into <c>recordVideo.dir</c> directory. If
		/// not specified videos are not recorded. Make sure to await <see cref="BrowserContext.Close"/>
		/// for videos to be saved.
		/// </param>
		/// <param name="timezoneId">
		/// Changes the timezone of the context. See <a href="https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1">ICU's
		/// metaZones.txt</a> for a list of supported timezone IDs.
		/// </param>
		/// <param name="userAgent">Specific user agent to use in this context.</param>
		/// <param name="viewport">
		/// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. <c>no_viewport</c>
		/// disables the fixed viewport.
		/// </param>
		NewContext(*bool acceptDownloads = default, *bool bypassCSP = default, *ColorScheme colorScheme = default, *float64 deviceScaleFactor = default, map[string]string extraHTTPHeaders = default, *BrowserGeolocation geolocation = default, *bool hasTouch = default, *BrowserHttpCredentials httpCredentials = default, *bool ignoreHTTPSErrors = default, *bool isMobile = default, *bool javaScriptEnabled = default, *string locale = default, *bool offline = default, []string permissions = default, *BrowserProxy proxy = default, *BrowserRecordVideo recordVideo = default, *string timezoneId = default, *string userAgent = default, *BrowserViewport viewport = default) BrowserContext
		/// <summary>
		/// <para>
		/// Creates a new page in a new browser context. Closing this page will close the context
		/// as well.
		/// </para>
		/// <para>
		/// This is a convenience API that should only be used for the single-page scenarios
		/// and short snippets. Production code and testing frameworks should explicitly create
		/// <see cref="Browser.NewContext"/> followed by the <see cref="BrowserContext.NewPage"/>
		/// to control their exact life times.
		/// </para>
		/// </summary>
		/// <param name="acceptDownloads">
		/// Whether to automatically download all the attachments. Defaults to <c>false</c>
		/// where all the downloads are canceled.
		/// </param>
		/// <param name="bypassCSP">Toggles bypassing page's Content-Security-Policy.</param>
		/// <param name="colorScheme">
		/// Emulates <c>'prefers-colors-scheme'</c> media feature, supported values are <c>'light'</c>,
		/// <c>'dark'</c>, <c>'no-preference'</c>. See <see cref="Page.EmulateMedia"/> for more
		/// details. Defaults to <c>'light'</c>.
		/// </param>
		/// <param name="deviceScaleFactor">Specify device scale factor (can be thought of as dpr). Defaults to <c>1</c>.</param>
		/// <param name="extraHTTPHeaders">
		/// An object containing additional HTTP headers to be sent with every request. All
		/// header values must be strings.
		/// </param>
		/// <param name="geolocation">
		/// </param>
		/// <param name="hasTouch">Specifies if viewport supports touch events. Defaults to false.</param>
		/// <param name="httpCredentials">
		/// Credentials for <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication">HTTP
		/// authentication</a>.
		/// </param>
		/// <param name="ignoreHTTPSErrors">Whether to ignore HTTPS errors during navigation. Defaults to <c>false</c>.</param>
		/// <param name="isMobile">
		/// Whether the <c>meta viewport</c> tag is taken into account and touch events are
		/// enabled. Defaults to <c>false</c>. Not supported in Firefox.
		/// </param>
		/// <param name="javaScriptEnabled">Whether or not to enable JavaScript in the context. Defaults to <c>true</c>.</param>
		/// <param name="locale">
		/// Specify user locale, for example <c>en-GB</c>, <c>de-DE</c>, etc. Locale will affect
		/// <c>navigator.language</c> value, <c>Accept-Language</c> request header value as
		/// well as number and date formatting rules.
		/// </param>
		/// <param name="offline">Whether to emulate network being offline. Defaults to <c>false</c>.</param>
		/// <param name="permissions">
		/// A list of permissions to grant to all pages in this context. See <see cref="BrowserContext.GrantPermissions"/>
		/// for more details.
		/// </param>
		/// <param name="proxy">
		/// Network proxy settings to use with this context. Note that browser needs to be launched
		/// with the global proxy for this option to work. If all contexts override the proxy,
		/// global proxy will be never used and can be any string, for example <c>launch({ proxy:
		/// { server: 'per-context' } })</c>.
		/// </param>
		/// <param name="recordVideo">
		/// Enables video recording for all pages into <c>recordVideo.dir</c> directory. If
		/// not specified videos are not recorded. Make sure to await <see cref="BrowserContext.Close"/>
		/// for videos to be saved.
		/// </param>
		/// <param name="timezoneId">
		/// Changes the timezone of the context. See <a href="https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1">ICU's
		/// metaZones.txt</a> for a list of supported timezone IDs.
		/// </param>
		/// <param name="userAgent">Specific user agent to use in this context.</param>
		/// <param name="viewport">
		/// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. <c>no_viewport</c>
		/// disables the fixed viewport.
		/// </param>
		NewPage(*bool acceptDownloads = default, *bool bypassCSP = default, *ColorScheme colorScheme = default, *float64 deviceScaleFactor = default, map[string]string extraHTTPHeaders = default, *BrowserGeolocation geolocation = default, *bool hasTouch = default, *BrowserHttpCredentials httpCredentials = default, *bool ignoreHTTPSErrors = default, *bool isMobile = default, *bool javaScriptEnabled = default, *string locale = default, *bool offline = default, []string permissions = default, *BrowserProxy proxy = default, *BrowserRecordVideo recordVideo = default, *string timezoneId = default, *string userAgent = default, *BrowserViewport viewport = default) Page
		/// <summary><para>Returns the browser version.</para></summary>
		*string Version { get; }
	}/// <summary>
	/// <para>BrowserContexts provide a way to operate multiple independent browser sessions.</para>
	/// <para>
	/// If a page opens another page, e.g. with a <c>window.open</c> call, the popup will
	/// belong to the parent page's browser context.
	/// </para>
	/// <para>
	/// Playwright allows creation of "incognito" browser contexts with <c>browser.newContext()</c>
	/// method. "Incognito" browser contexts don't write any browsing data to disk.
	/// </para>
	/// </summary>
	type BrowserContext interface {
	EventEmitter
		/// <summary>
		/// <para>
		/// Emitted when Browser context gets closed. This might happen because of one of the
		/// following:
		/// </para>
		/// <list type="bullet">
		/// <item><description>Browser context is closed.</description></item>
		/// <item><description>Browser application is closed or crashed.</description></item>
		/// <item><description>The <see cref="Browser.Close"/> method was called.</description></item>
		/// </list>
		/// </summary>
		event EventHandler Close;
		/// <summary>
		/// <para>
		/// The event is emitted when a new Page is created in the BrowserContext. The page
		/// may still be loading. The event will also fire for popup pages. See also <see cref="Page.Popup"/>
		/// to receive events about popups relevant to a specific page.
		/// </para>
		/// <para>
		/// The earliest moment that page is available is when it has navigated to the initial
		/// url. For example, when opening a popup with <c>window.open('http://example.com')</c>,
		/// this event will fire when the network request to "http://example.com" is done and
		/// its response has started loading in the popup.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Use <see cref="Page.WaitForLoadState"/> to wait until the page gets to a particular
		/// state (you should not need it in most cases).
		/// </para>
		/// </remarks>
		event EventHandler<Page> Page;
		/// <summary>
		/// <para>
		/// Adds cookies into this browser context. All pages within this context will have
		/// these cookies installed. Cookies can be obtained via <see cref="BrowserContext.Cookies"/>.
		/// </para>
		/// </summary>
		/// <param name="cookies">
		/// </param>
		AddCookies([]BrowserContextCookies cookies) void
		/// <summary>
		/// <para>Adds a script which would be evaluated in one of the following scenarios:</para>
		/// <list type="bullet">
		/// <item><description>Whenever a page is created in the browser context or is navigated.</description></item>
		/// <item><description>
		/// Whenever a child frame is attached or navigated in any page in the browser context.
		/// In this case, the script is evaluated in the context of the newly attached frame.
		/// </description></item>
		/// </list>
		/// <para>
		/// The script is evaluated after the document was created but before any of its scripts
		/// were run. This is useful to amend the JavaScript environment, e.g. to seed <c>Math.random</c>.
		/// </para>
		/// <para>An example of overriding <c>Math.random</c> before the page loads:</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// The order of evaluation of multiple scripts installed via <see cref="BrowserContext.AddInitScript"/>
		/// and <see cref="Page.AddInitScript"/> is not defined.
		/// </para>
		/// </remarks>
		/// <param name="script">Optional Script source to be evaluated in all pages in the browser context.</param>
		/// <param name="path">Optional Script path to be evaluated in all pages in the browser context.</param>
		AddInitScript(*string script = default, *string path = default) void
		/// <summary>
		/// <para>
		/// Returns the browser instance of the context. If it was launched as a persistent
		/// context null gets returned.
		/// </para>
		/// </summary>
		Browser Browser { get; }
		/// <summary><para>Clears context cookies.</para></summary>
		ClearCookies() void
		/// <summary><para>Clears all permission overrides for the browser context.</para></summary>
		ClearPermissions() void
		/// <summary>
		/// <para>
		/// Closes the browser context. All the pages that belong to the browser context will
		/// be closed.
		/// </para>
		/// </summary>
		/// <remarks><para>The default browser context cannot be closed.</para></remarks>
		Close() void
		/// <summary>
		/// <para>
		/// If no URLs are specified, this method returns all cookies. If URLs are specified,
		/// only cookies that affect those URLs are returned.
		/// </para>
		/// </summary>
		/// <param name="urls">Optional list of URLs.</param>
		Cookies([]string urls = default) IReadOnlyCollection<*BrowserContextCookiesResult>
		/// <summary>
		/// <para>
		/// The method adds a function called <paramref name="name"/> on the <c>window</c> object
		/// of every frame in every page in the context. When called, the function executes
		/// <paramref name="callback"/> and returns a <see cref="Promise"/> which resolves to
		/// the return value of <paramref name="callback"/>. If the <paramref name="callback"/>
		/// returns a <see cref="Promise"/>, it will be awaited.
		/// </para>
		/// <para>
		/// The first argument of the <paramref name="callback"/> function contains information
		/// about the caller: <c>{ browserContext: BrowserContext, page: Page, frame: Frame
		/// }</c>.
		/// </para>
		/// <para>See <see cref="Page.ExposeBinding"/> for page-only version.</para>
		/// <para>An example of exposing page URL to all frames in all pages in the context:</para>
		/// <para>An example of passing an element handle:</para>
		/// </summary>
		/// <param name="name">Name of the function on the window object.</param>
		/// <param name="callback">Callback function that will be called in the Playwright's context.</param>
		/// <param name="handle">
		/// Whether to pass the argument as a handle, instead of passing by value. When passing
		/// a handle, only one argument is supported. When passing by value, multiple arguments
		/// are supported.
		/// </param>
		ExposeBinding(*string name, interface{} callback, *bool handle = default) void
		/// <summary>
		/// <para>
		/// The method adds a function called <paramref name="name"/> on the <c>window</c> object
		/// of every frame in every page in the context. When called, the function executes
		/// <paramref name="callback"/> and returns a <see cref="Promise"/> which resolves to
		/// the return value of <paramref name="callback"/>.
		/// </para>
		/// <para>If the <paramref name="callback"/> returns a <see cref="Promise"/>, it will be awaited.</para>
		/// <para>See <see cref="Page.ExposeFunction"/> for page-only version.</para>
		/// <para>An example of adding an <c>md5</c> function to all pages in the context:</para>
		/// </summary>
		/// <param name="name">Name of the function on the window object.</param>
		/// <param name="callback">Callback function that will be called in the Playwright's context.</param>
		ExposeFunction(*string name, interface{} callback) void
		/// <summary>
		/// <para>
		/// Grants specified permissions to the browser context. Only grants corresponding permissions
		/// to the given origin if specified.
		/// </para>
		/// </summary>
		/// <param name="permissions">
		/// A permission or an array of permissions to grant. Permissions can be one of the
		/// following values:
		/// <list type="bullet">
		/// <item><description><c>'geolocation'</c></description></item>
		/// <item><description><c>'midi'</c></description></item>
		/// <item><description><c>'midi-sysex'</c> (system-exclusive midi)</description></item>
		/// <item><description><c>'notifications'</c></description></item>
		/// <item><description><c>'push'</c></description></item>
		/// <item><description><c>'camera'</c></description></item>
		/// <item><description><c>'microphone'</c></description></item>
		/// <item><description><c>'background-sync'</c></description></item>
		/// <item><description><c>'ambient-light-sensor'</c></description></item>
		/// <item><description><c>'accelerometer'</c></description></item>
		/// <item><description><c>'gyroscope'</c></description></item>
		/// <item><description><c>'magnetometer'</c></description></item>
		/// <item><description><c>'accessibility-events'</c></description></item>
		/// <item><description><c>'clipboard-read'</c></description></item>
		/// <item><description><c>'clipboard-write'</c></description></item>
		/// <item><description><c>'payment-handler'</c></description></item>
		/// </list>
		/// </param>
		/// <param name="origin">The <see cref="origin"/> to grant permissions to, e.g. "https://example.com".</param>
		GrantPermissions([]string permissions, *string origin = default) void
		/// <summary><para>Creates a new page in the browser context.</para></summary>
		GetNewPage() Page
		/// <summary><para>Returns all open pages in the context.</para></summary>
		IReadOnlyCollection<Page> Pages { get; }
		/// <summary>
		/// <para>
		/// Routing provides the capability to modify network requests that are made by any
		/// page in the browser context. Once route is enabled, every request matching the url
		/// pattern will stall unless it's continued, fulfilled or aborted.
		/// </para>
		/// <para>An example of a naïve handler that aborts all image requests:</para>
		/// <para>or the same snippet using a regex pattern instead:</para>
		/// <para>
		/// Page routes (set up with <see cref="Page.Route"/>) take precedence over browser
		/// context routes when request matches both handlers.
		/// </para>
		/// </summary>
		/// <remarks><para>Enabling routing disables http cache.</para></remarks>
		/// <param name="url">
		/// A glob pattern, regex pattern or predicate receiving <see cref="URL"/> to match
		/// while routing.
		/// </param>
		/// <param name="handler">handler function to route the request.</param>
		Route(interface{} url, func(Route) handler) void
		/// <summary>
		/// <para>
		/// This setting will change the default maximum navigation time for the following methods
		/// and related shortcuts:
		/// </para>
		/// <list type="bullet">
		/// <item><description><see cref="Page.GoBack"/></description></item>
		/// <item><description><see cref="Page.GoForward"/></description></item>
		/// <item><description><see cref="Page.Goto"/></description></item>
		/// <item><description><see cref="Page.Reload"/></description></item>
		/// <item><description><see cref="Page.SetContent"/></description></item>
		/// <item><description><see cref="Page.WaitForNavigation"/></description></item>
		/// </list>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <see cref="Page.SetDefaultNavigationTimeout"/> and <see cref="Page.SetDefaultTimeout"/>
		/// take priority over <see cref="BrowserContext.SetDefaultNavigationTimeout"/>.
		/// </para>
		/// </remarks>
		*float64 DefaultNavigationTimeout { set; }
		/// <summary>
		/// <para>
		/// This setting will change the default maximum time for all the methods accepting
		/// <paramref name="timeout"/> option.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <see cref="Page.SetDefaultNavigationTimeout"/>, <see cref="Page.SetDefaultTimeout"/>
		/// and <see cref="BrowserContext.SetDefaultNavigationTimeout"/> take priority over
		/// <see cref="BrowserContext.SetDefaultTimeout"/>.
		/// </para>
		/// </remarks>
		*float64 DefaultTimeout { set; }
		/// <summary>
		/// <para>
		/// The extra HTTP headers will be sent with every request initiated by any page in
		/// the context. These headers are merged with page-specific extra HTTP headers set
		/// with <see cref="Page.SetExtraHttpHeaders"/>. If page overrides a particular header,
		/// page-specific header value will be used instead of the browser context header value.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <see cref="BrowserContext.SetExtraHttpHeaders"/> does not guarantee the order of
		/// headers in the outgoing requests.
		/// </para>
		/// </remarks>
		/// <param name="headers">
		/// An object containing additional HTTP headers to be sent with every request. All
		/// header values must be strings.
		/// </param>
		SetExtraHttpHeaders(map[string]string headers) void
		/// <summary>
		/// <para>
		/// Sets the context's geolocation. Passing <c>null</c> or <c>undefined</c> emulates
		/// position unavailable.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Consider using <see cref="BrowserContext.GrantPermissions"/> to grant permissions
		/// for the browser context pages to read its geolocation.
		/// </para>
		/// </remarks>
		/// <param name="geolocation">
		/// </param>
		SetGeolocation(*BrowserContextGeolocation geolocation) void
		/// <param name="offline">Whether to emulate network being offline for the browser context.</param>
		SetOffline(*bool offline) void
		/// <summary>
		/// <para>
		/// Returns storage state for this browser context, contains current cookies and local
		/// storage snapshot.
		/// </para>
		/// </summary>
		/// <param name="path">
		/// The file path to save the storage state to. If <paramref name="path"/> is a relative
		/// path, then it is resolved relative to current working directory. If no path is provided,
		/// storage state is still returned, but won't be saved to the disk.
		/// </param>
		StorageState(*string path = default) *BrowserContextStorageStateResult
		/// <summary>
		/// <para>
		/// Removes a route created with <see cref="BrowserContext.Route"/>. When <paramref
		/// name="handler"/> is not specified, removes all routes for the <paramref name="url"/>.
		/// </para>
		/// </summary>
		/// <param name="url">
		/// A glob pattern, regex pattern or predicate receiving <see cref="URL"/> used to register
		/// a routing with <see cref="BrowserContext.Route"/>.
		/// </param>
		/// <param name="handler">Optional handler function used to register a routing with <see cref="BrowserContext.Route"/>.</param>
		Unroute(interface{} url, func(Route, Request) handler = default) void
	}/// <summary>
	/// <para>
	/// BrowserType provides methods to launch a specific browser instance or connect to
	/// an existing one. The following is a typical example of using Playwright to drive
	/// automation:
	/// </para>
	/// </summary>
	type BrowserType interface {
		/// <summary><para>A path where Playwright expects to find a bundled browser executable.</para></summary>
		*string ExecutablePath { get; }
		/// <summary>
		/// <para>Returns the browser instance.</para>
		/// <para>
		/// You can use <paramref name="ignoreDefaultArgs"/> to filter out <c>--mute-audio</c>
		/// from default arguments:
		/// </para>
		/// <para>
		/// > **Chromium-only** Playwright can also be used to control the Google Chrome or
		/// Microsoft Edge browsers, but it works best with the version of Chromium it is bundled
		/// with. There is no guarantee it will work with any other version. Use <paramref name="executablePath"/>
		/// option with extreme caution.
		/// </para>
		/// <para>></para>
		/// <para>
		/// > If Google Chrome (rather than Chromium) is preferred, a <a href="https://www.google.com/chrome/browser/canary.html">Chrome
		/// Canary</a> or <a href="https://www.chromium.org/getting-involved/dev-channel">Dev
		/// Channel</a> build is suggested.
		/// </para>
		/// <para>></para>
		/// <para>
		/// > Stock browsers like Google Chrome and Microsoft Edge are suitable for tests that
		/// require proprietary media codecs for video playback. See <a href="https://www.howtogeek.com/202825/what%E2%80%99s-the-difference-between-chromium-and-chrome/">this
		/// article</a> for other differences between Chromium and Chrome. <a href="https://chromium.googlesource.com/chromium/src/+/lkgr/docs/chromium_browser_vs_google_chrome.md">This
		/// article</a> describes some differences for Linux users.
		/// </para>
		/// </summary>
		/// <param name="args">
		/// Additional arguments to pass to the browser instance. The list of Chromium flags
		/// can be found <a href="http://peter.sh/experiments/chromium-command-line-switches/">here</a>.
		/// </param>
		/// <param name="chromiumSandbox">Enable Chromium sandboxing. Defaults to <c>false</c>.</param>
		/// <param name="devtools">
		/// **Chromium-only** Whether to auto-open a Developer Tools panel for each tab. If
		/// this option is <c>true</c>, the <paramref name="headless"/> option will be set <c>false</c>.
		/// </param>
		/// <param name="downloadsPath">
		/// If specified, accepted downloads are downloaded into this directory. Otherwise,
		/// temporary directory is created and is deleted when browser is closed.
		/// </param>
		/// <param name="executablePath">
		/// Path to a browser executable to run instead of the bundled one. If <paramref name="executablePath"/>
		/// is a relative path, then it is resolved relative to the current working directory.
		/// Note that Playwright only works with the bundled Chromium, Firefox or WebKit, use
		/// at your own risk.
		/// </param>
		/// <param name="handleSIGHUP">Close the browser process on SIGHUP. Defaults to <c>true</c>.</param>
		/// <param name="handleSIGINT">Close the browser process on Ctrl-C. Defaults to <c>true</c>.</param>
		/// <param name="handleSIGTERM">Close the browser process on SIGTERM. Defaults to <c>true</c>.</param>
		/// <param name="headless">
		/// Whether to run browser in headless mode. More details for <a href="https://developers.google.com/web/updates/2017/04/headless-chrome">Chromium</a>
		/// and <a href="https://developer.mozilla.org/en-US/docs/Mozilla/Firefox/Headless_mode">Firefox</a>.
		/// Defaults to <c>true</c> unless the <paramref name="devtools"/> option is <c>true</c>.
		/// </param>
		/// <param name="proxy">Network proxy settings.</param>
		/// <param name="slowMo">
		/// Slows down Playwright operations by the specified amount of milliseconds. Useful
		/// so that you can see what is going on.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds to wait for the browser instance to start. Defaults
		/// to <c>30000</c> (30 seconds). Pass <c>0</c> to disable timeout.
		/// </param>
		Launch([]string args = default, *bool chromiumSandbox = default, *bool devtools = default, *string downloadsPath = default, *string executablePath = default, *bool handleSIGHUP = default, *bool handleSIGINT = default, *bool handleSIGTERM = default, *bool headless = default, *BrowserTypeProxy proxy = default, *float64 slowMo = default, *float64 timeout = default) Browser
		/// <summary>
		/// <para>Returns the persistent browser context instance.</para>
		/// <para>
		/// Launches browser that uses persistent storage located at <paramref name="userDataDir"/>
		/// and returns the only context. Closing this context will automatically close the
		/// browser.
		/// </para>
		/// </summary>
		/// <param name="userDataDir">
		/// Path to a User Data Directory, which stores browser session data like cookies and
		/// local storage. More details for <a href="https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md#introduction">Chromium</a>
		/// and <a href="https://developer.mozilla.org/en-US/docs/Mozilla/Command_Line_Options#User_Profile">Firefox</a>.
		/// Note that Chromium's user data directory is the **parent** directory of the "Profile
		/// Path" seen at <c>chrome://version</c>.
		/// </param>
		/// <param name="acceptDownloads">
		/// Whether to automatically download all the attachments. Defaults to <c>false</c>
		/// where all the downloads are canceled.
		/// </param>
		/// <param name="args">
		/// Additional arguments to pass to the browser instance. The list of Chromium flags
		/// can be found <a href="http://peter.sh/experiments/chromium-command-line-switches/">here</a>.
		/// </param>
		/// <param name="bypassCSP">Toggles bypassing page's Content-Security-Policy.</param>
		/// <param name="chromiumSandbox">Enable Chromium sandboxing. Defaults to <c>true</c>.</param>
		/// <param name="colorScheme">
		/// Emulates <c>'prefers-colors-scheme'</c> media feature, supported values are <c>'light'</c>,
		/// <c>'dark'</c>, <c>'no-preference'</c>. See <see cref="Page.EmulateMedia"/> for more
		/// details. Defaults to <c>'light'</c>.
		/// </param>
		/// <param name="deviceScaleFactor">Specify device scale factor (can be thought of as dpr). Defaults to <c>1</c>.</param>
		/// <param name="devtools">
		/// **Chromium-only** Whether to auto-open a Developer Tools panel for each tab. If
		/// this option is <c>true</c>, the <paramref name="headless"/> option will be set <c>false</c>.
		/// </param>
		/// <param name="downloadsPath">
		/// If specified, accepted downloads are downloaded into this directory. Otherwise,
		/// temporary directory is created and is deleted when browser is closed.
		/// </param>
		/// <param name="executablePath">
		/// Path to a browser executable to run instead of the bundled one. If <paramref name="executablePath"/>
		/// is a relative path, then it is resolved relative to the current working directory.
		/// **BEWARE**: Playwright is only guaranteed to work with the bundled Chromium, Firefox
		/// or WebKit, use at your own risk.
		/// </param>
		/// <param name="extraHTTPHeaders">
		/// An object containing additional HTTP headers to be sent with every request. All
		/// header values must be strings.
		/// </param>
		/// <param name="geolocation">
		/// </param>
		/// <param name="handleSIGHUP">Close the browser process on SIGHUP. Defaults to <c>true</c>.</param>
		/// <param name="handleSIGINT">Close the browser process on Ctrl-C. Defaults to <c>true</c>.</param>
		/// <param name="handleSIGTERM">Close the browser process on SIGTERM. Defaults to <c>true</c>.</param>
		/// <param name="hasTouch">Specifies if viewport supports touch events. Defaults to false.</param>
		/// <param name="headless">
		/// Whether to run browser in headless mode. More details for <a href="https://developers.google.com/web/updates/2017/04/headless-chrome">Chromium</a>
		/// and <a href="https://developer.mozilla.org/en-US/docs/Mozilla/Firefox/Headless_mode">Firefox</a>.
		/// Defaults to <c>true</c> unless the <paramref name="devtools"/> option is <c>true</c>.
		/// </param>
		/// <param name="httpCredentials">
		/// Credentials for <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication">HTTP
		/// authentication</a>.
		/// </param>
		/// <param name="ignoreHTTPSErrors">Whether to ignore HTTPS errors during navigation. Defaults to <c>false</c>.</param>
		/// <param name="isMobile">
		/// Whether the <c>meta viewport</c> tag is taken into account and touch events are
		/// enabled. Defaults to <c>false</c>. Not supported in Firefox.
		/// </param>
		/// <param name="javaScriptEnabled">Whether or not to enable JavaScript in the context. Defaults to <c>true</c>.</param>
		/// <param name="locale">
		/// Specify user locale, for example <c>en-GB</c>, <c>de-DE</c>, etc. Locale will affect
		/// <c>navigator.language</c> value, <c>Accept-Language</c> request header value as
		/// well as number and date formatting rules.
		/// </param>
		/// <param name="offline">Whether to emulate network being offline. Defaults to <c>false</c>.</param>
		/// <param name="permissions">
		/// A list of permissions to grant to all pages in this context. See <see cref="BrowserContext.GrantPermissions"/>
		/// for more details.
		/// </param>
		/// <param name="proxy">Network proxy settings.</param>
		/// <param name="recordVideo">
		/// Enables video recording for all pages into <c>recordVideo.dir</c> directory. If
		/// not specified videos are not recorded. Make sure to await <see cref="BrowserContext.Close"/>
		/// for videos to be saved.
		/// </param>
		/// <param name="slowMo">
		/// Slows down Playwright operations by the specified amount of milliseconds. Useful
		/// so that you can see what is going on. Defaults to 0.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds to wait for the browser instance to start. Defaults
		/// to <c>30000</c> (30 seconds). Pass <c>0</c> to disable timeout.
		/// </param>
		/// <param name="timezoneId">
		/// Changes the timezone of the context. See <a href="https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1">ICU's
		/// metaZones.txt</a> for a list of supported timezone IDs.
		/// </param>
		/// <param name="userAgent">Specific user agent to use in this context.</param>
		/// <param name="viewport">
		/// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. <c>no_viewport</c>
		/// disables the fixed viewport.
		/// </param>
		LaunchPersistentContext(*string userDataDir, *bool acceptDownloads = default, []string args = default, *bool bypassCSP = default, *bool chromiumSandbox = default, *ColorScheme colorScheme = default, *float64 deviceScaleFactor = default, *bool devtools = default, *string downloadsPath = default, *string executablePath = default, map[string]string extraHTTPHeaders = default, *BrowserTypeGeolocation geolocation = default, *bool handleSIGHUP = default, *bool handleSIGINT = default, *bool handleSIGTERM = default, *bool hasTouch = default, *bool headless = default, *BrowserTypeHttpCredentials httpCredentials = default, *bool ignoreHTTPSErrors = default, *bool isMobile = default, *bool javaScriptEnabled = default, *string locale = default, *bool offline = default, []string permissions = default, *BrowserTypeProxy proxy = default, *BrowserTypeRecordVideo recordVideo = default, *float64 slowMo = default, *float64 timeout = default, *string timezoneId = default, *string userAgent = default, *BrowserTypeViewport viewport = default) BrowserContext
		/// <summary><para>Returns browser name. For example: <c>'chromium'</c>, <c>'webkit'</c> or <c>'firefox'</c>.</para></summary>
		*string Name { get; }
	}/// <summary>
	/// <para>
	/// <see cref="ConsoleMessage"/> objects are dispatched by page via the <see cref="Page.Console"/>
	/// event.
	/// </para>
	/// </summary>
	type ConsoleMessage interface {
		IReadOnlyCollection<JSHandle> Args { get; }
		*string Text { get; }
		/// <summary>
		/// <para>
		/// One of the following values: <c>'log'</c>, <c>'debug'</c>, <c>'info'</c>, <c>'error'</c>,
		/// <c>'warning'</c>, <c>'dir'</c>, <c>'dirxml'</c>, <c>'table'</c>, <c>'trace'</c>,
		/// <c>'clear'</c>, <c>'startGroup'</c>, <c>'startGroupCollapsed'</c>, <c>'endGroup'</c>,
		/// <c>'assert'</c>, <c>'profile'</c>, <c>'profileEnd'</c>, <c>'count'</c>, <c>'timeEnd'</c>.
		/// </para>
		/// </summary>
		*string Type { get; }
	}/// <summary>
	/// <para>
	/// <see cref="Dialog"/> objects are dispatched by page via the <see cref="Page.Dialog"/>
	/// event.
	/// </para>
	/// <para>An example of using <c>Dialog</c> class:</para>
	/// </summary>
	/// <remarks>
	/// <para>
	/// Dialogs are dismissed automatically, unless there is a <see cref="Page.Dialog"/>
	/// listener. When listener is present, it **must** either <see cref="Dialog.Accept"/>
	/// or <see cref="Dialog.Dismiss"/> the dialog - otherwise the page will <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking">freeze</a>
	/// waiting for the dialog, and actions like click will never finish.
	/// </para>
	/// </remarks>
	type Dialog interface {
		/// <summary><para>Returns when the dialog has been accepted.</para></summary>
		/// <param name="promptText">
		/// A text to enter in prompt. Does not cause any effects if the dialog's <c>type</c>
		/// is not prompt. Optional.
		/// </param>
		Accept(*string promptText = default) void
		/// <summary><para>If dialog is prompt, returns default prompt value. Otherwise, returns empty string.</para></summary>
		*string DefaultValue { get; }
		/// <summary><para>Returns when the dialog has been dismissed.</para></summary>
		Dismiss() void
		/// <summary><para>A message displayed in the dialog.</para></summary>
		*string Message { get; }
		/// <summary>
		/// <para>
		/// Returns dialog's type, can be one of <c>alert</c>, <c>beforeunload</c>, <c>confirm</c>
		/// or <c>prompt</c>.
		/// </para>
		/// </summary>
		*string Type { get; }
	}/// <summary>
	/// <para>
	/// <see cref="Download"/> objects are dispatched by page via the <see cref="Page.Download"/>
	/// event.
	/// </para>
	/// <para>
	/// All the downloaded files belonging to the browser context are deleted when the browser
	/// context is closed. All downloaded files are deleted when the browser closes.
	/// </para>
	/// <para>
	/// Download event is emitted once the download starts. Download path becomes available
	/// once download completes:
	/// </para>
	/// </summary>
	/// <remarks>
	/// <para>
	/// Browser context **must** be created with the <paramref name="acceptDownloads"/>
	/// set to <c>true</c> when user needs access to the downloaded content. If <paramref
	/// name="acceptDownloads"/> is not set, download events are emitted, but the actual
	/// download is not performed and user has no access to the downloaded files.
	/// </para>
	/// </remarks>
	type Download interface {
		/// <summary><para>Deletes the downloaded file. Will wait for the download to finish if necessary.</para></summary>
		Delete() void
		/// <summary><para>Returns download error if any. Will wait for the download to finish if necessary.</para></summary>
		GetFailure() *string
		/// <summary>
		/// <para>
		/// Returns path to the downloaded file in case of successful download. The method will
		/// wait for the download to finish if necessary.
		/// </para>
		/// </summary>
		GetPath() *string
		/// <summary>
		/// <para>
		/// Saves the download to a user-specified path. It is safe to call this method while
		/// the download is still in progress.
		/// </para>
		/// </summary>
		/// <param name="path">Path where the download should be saved.</param>
		SaveAs(*string path) void
		/// <summary>
		/// <para>
		/// Returns suggested filename for this download. It is typically computed by the browser
		/// from the <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition"><c>Content-Disposition</c></a>
		/// response header or the <c>download</c> attribute. See the spec on <a href="https://html.spec.whatwg.org/#downloading-resources">whatwg</a>.
		/// Different browsers can use different logic for computing it.
		/// </para>
		/// </summary>
		*string SuggestedFilename { get; }
		/// <summary><para>Returns downloaded url.</para></summary>
		*string URL { get; }
	}/// <summary>
	/// <para>
	/// ElementHandle represents an in-page DOM element. ElementHandles can be created with
	/// the <see cref="Page.QuerySelector"/> method.
	/// </para>
	/// <para>
	/// ElementHandle prevents DOM element from garbage collection unless the handle is
	/// disposed with <see cref="JSHandle.Dispose"/>. ElementHandles are auto-disposed when
	/// their origin frame gets navigated.
	/// </para>
	/// <para>
	/// ElementHandle instances can be used as an argument in <see cref="Page.EvalOnSelector"/>
	/// and <see cref="Page.Evaluate"/> methods.
	/// </para>
	/// </summary>
	/// <remarks>Inherits from <see cref="JSHandle"/></remarks>
	type ElementHandle interface {
	JSHandle
		/// <summary>
		/// <para>
		/// This method returns the bounding box of the element, or <c>null</c> if the element
		/// is not visible. The bounding box is calculated relative to the main frame viewport
		/// - which is usually the same as the browser window.
		/// </para>
		/// <para>
		/// Scrolling affects the returned bonding box, similarly to <a href="https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect">Element.getBoundingClientRect</a>.
		/// That means <c>x</c> and/or <c>y</c> may be negative.
		/// </para>
		/// <para>
		/// Elements from child frames return the bounding box relative to the main frame, unlike
		/// the <a href="https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect">Element.getBoundingClientRect</a>.
		/// </para>
		/// <para>
		/// Assuming the page is static, it is safe to use bounding box coordinates to perform
		/// input. For example, the following snippet should click the center of the element.
		/// </para>
		/// </summary>
		GetBoundingBox() *ElementHandleBoundingBoxResult
		/// <summary>
		/// <para>This method checks the element by performing the following steps:</para>
		/// <list type="ordinal">
		/// <item><description>
		/// Ensure that element is a checkbox or a radio input. If not, this method rejects.
		/// If the element is already checked, this method returns immediately.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the element, unless
		/// <paramref name="force"/> option is set.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>Use <see cref="Page.Mouse"/> to click in the center of the element.</description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// <item><description>Ensure that the element is now checked. If not, this method rejects.</description></item>
		/// </list>
		/// <para>
		/// If the element is detached from the DOM at any moment during the action, this method
		/// rejects.
		/// </para>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Check(*bool force = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>This method clicks the element by performing the following steps:</para>
		/// <list type="ordinal">
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the element, unless
		/// <paramref name="force"/> option is set.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to click in the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// If the element is detached from the DOM at any moment during the action, this method
		/// rejects.
		/// </para>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="clickCount">defaults to 1. See <see cref="UIEvent.detail"/>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Click(*MouseButton button = default, *int clickCount = default, *float64 delay = default, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *ElementHandlePosition position = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Returns the content frame for element handles referencing iframe nodes, or <c>null</c>
		/// otherwise
		/// </para>
		/// </summary>
		GetContentFrame() Frame
		/// <summary>
		/// <para>This method double clicks the element by performing the following steps:</para>
		/// <list type="ordinal">
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the element, unless
		/// <paramref name="force"/> option is set.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to double click in the center of the element, or the
		/// specified <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set. Note that if the first click of the <c>dblclick()</c> triggers a
		/// navigation event, this method will reject.
		/// </description></item>
		/// </list>
		/// <para>
		/// If the element is detached from the DOM at any moment during the action, this method
		/// rejects.
		/// </para>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>elementHandle.dblclick()</c> dispatches two <c>click</c> events and a single
		/// <c>dblclick</c> event.
		/// </para>
		/// </remarks>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Dblclick(*MouseButton button = default, *float64 delay = default, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *ElementHandlePosition position = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// The snippet below dispatches the <c>click</c> event on the element. Regardless of
		/// the visibility state of the elment, <c>click</c> is dispatched. This is equivalend
		/// to calling <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click">element.click()</a>.
		/// </para>
		/// <para>
		/// Under the hood, it creates an instance of an event based on the given <paramref
		/// name="type"/>, initializes it with <paramref name="eventInit"/> properties and dispatches
		/// it on the element. Events are <c>composed</c>, <c>cancelable</c> and bubble by default.
		/// </para>
		/// <para>
		/// Since <paramref name="eventInit"/> is event-specific, please refer to the events
		/// documentation for the lists of initial properties:
		/// </para>
		/// <list type="bullet">
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent">DragEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent">FocusEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent">KeyboardEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent">MouseEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent">PointerEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent">TouchEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/Event/Event">Event</a></description></item>
		/// </list>
		/// <para>
		/// You can also specify <c>JSHandle</c> as the property value if you want live objects
		/// to be passed into the event:
		/// </para>
		/// </summary>
		/// <param name="type">DOM event type: <c>"click"</c>, <c>"dragstart"</c>, etc.</param>
		/// <param name="eventInit">Optional event-specific initialization properties.</param>
		DispatchEvent(*string type, interface{} eventInit = default) void
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>
		/// The method finds an element matching the specified selector in the <c>ElementHandle</c>s
		/// subtree and passes it as a first argument to <paramref name="expression"/>. See
		/// <a href="./selectors.md">Working with selectors</a> for more details. If no elements
		/// match the selector, the method throws an error.
		/// </para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <see cref="ElementHandle.EvalOnSelector"/>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvalOnSelector(*string selector, *string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>
		/// The method finds all elements matching the specified selector in the <c>ElementHandle</c>'s
		/// subtree and passes an array of matched elements as a first argument to <paramref
		/// name="expression"/>. See <a href="./selectors.md">Working with selectors</a> for
		/// more details.
		/// </para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <see cref="ElementHandle.EvalOnSelectorAll"/>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvalOnSelectorAll(*string selector, *string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>
		/// This method waits for <a href="./actionability.md">actionability</a> checks, focuses
		/// the element, fills it and triggers an <c>input</c> event after filling. If the element
		/// is inside the <c>&lt;label&gt;</c> element that has associated <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control">control</a>,
		/// that control will be filled instead. If the element to be filled is not an <c>&lt;input&gt;</c>,
		/// <c>&lt;textarea&gt;</c> or <c>[contenteditable]</c> element, this method throws
		/// an error. Note that you can pass an empty string to clear the input field.
		/// </para>
		/// </summary>
		/// <param name="value">
		/// Value to set for the <c>&lt;input&gt;</c>, <c>&lt;textarea&gt;</c> or <c>[contenteditable]</c>
		/// element.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Fill(*string value, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Calls <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus">focus</a>
		/// on the element.
		/// </para>
		/// </summary>
		Focus() void
		/// <summary><para>Returns element attribute value.</para></summary>
		/// <param name="name">Attribute name to get the value for.</param>
		GetAttribute(*string name) *string
		/// <summary>
		/// <para>This method hovers over the element by performing the following steps:</para>
		/// <list type="ordinal">
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the element, unless
		/// <paramref name="force"/> option is set.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to hover over the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <c>noWaitAfter</c>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// If the element is detached from the DOM at any moment during the action, this method
		/// rejects.
		/// </para>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Hover(*bool force = default, []KeyboardModifier modifiers = default, *ElementHandlePosition position = default, *float64 timeout = default) void
		/// <summary><para>Returns the <c>element.innerHTML</c>.</para></summary>
		GetInnerHTML() *string
		/// <summary><para>Returns the <c>element.innerText</c>.</para></summary>
		GetInnerText() *string
		/// <summary>
		/// <para>
		/// Returns whether the element is checked. Throws if the element is not a checkbox
		/// or radio input.
		/// </para>
		/// </summary>
		GetIsChecked() *bool
		/// <summary><para>Returns whether the element is disabled, the opposite of <a href="./actionability.md#enabled">enabled</a>.</para></summary>
		GetIsDisabled() *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#editable">editable</a>.</para></summary>
		GetIsEditable() *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#enabled">enabled</a>.</para></summary>
		GetIsEnabled() *bool
		/// <summary><para>Returns whether the element is hidden, the opposite of <a href="./actionability.md#visible">visible</a>.</para></summary>
		GetIsHidden() *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#visible">visible</a>.</para></summary>
		GetIsVisible() *bool
		/// <summary><para>Returns the frame containing the given element.</para></summary>
		GetOwnerFrame() Frame
		/// <summary>
		/// <para>Focuses the element, and then uses <see cref="Keyboard.Down"/> and <see cref="Keyboard.Up"/>.</para>
		/// <para>
		/// <paramref name="key"/> can specify the intended <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key">keyboardEvent.key</a>
		/// value or a single character to generate the text for. A superset of the <paramref
		/// name="key"/> values can be found <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values">here</a>.
		/// Examples of the keys are:
		/// </para>
		/// <para>
		/// <c>F1</c> - <c>F12</c>, <c>Digit0</c>- <c>Digit9</c>, <c>KeyA</c>- <c>KeyZ</c>,
		/// <c>Backquote</c>, <c>Minus</c>, <c>Equal</c>, <c>Backslash</c>, <c>Backspace</c>,
		/// <c>Tab</c>, <c>Delete</c>, <c>Escape</c>, <c>ArrowDown</c>, <c>End</c>, <c>Enter</c>,
		/// <c>Home</c>, <c>Insert</c>, <c>PageDown</c>, <c>PageUp</c>, <c>ArrowRight</c>, <c>ArrowUp</c>,
		/// etc.
		/// </para>
		/// <para>
		/// Following modification shortcuts are also supported: <c>Shift</c>, <c>Control</c>,
		/// <c>Alt</c>, <c>Meta</c>, <c>ShiftLeft</c>.
		/// </para>
		/// <para>
		/// Holding down <c>Shift</c> will type the text that corresponds to the <paramref name="key"/>
		/// in the upper case.
		/// </para>
		/// <para>
		/// If <paramref name="key"/> is a single character, it is case-sensitive, so the values
		/// <c>a</c> and <c>A</c> will generate different respective texts.
		/// </para>
		/// <para>
		/// Shortcuts such as <c>key: "Control+o"</c> or <c>key: "Control+Shift+T"</c> are supported
		/// as well. When speficied with the modifier, modifier is pressed and being held while
		/// the subsequent key is being pressed.
		/// </para>
		/// </summary>
		/// <param name="key">
		/// Name of the key to press or a character to generate, such as <c>ArrowLeft</c> or
		/// <c>a</c>.
		/// </param>
		/// <param name="delay">
		/// Time to wait between <c>keydown</c> and <c>keyup</c> in milliseconds. Defaults to
		/// 0.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Press(*string key, *float64 delay = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// The method finds an element matching the specified selector in the <c>ElementHandle</c>'s
		/// subtree. See <a href="./selectors.md">Working with selectors</a> for more details.
		/// If no elements match the selector, returns <c>null</c>.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		QuerySelector(*string selector) ElementHandle
		/// <summary>
		/// <para>
		/// The method finds all elements matching the specified selector in the <c>ElementHandle</c>s
		/// subtree. See <a href="./selectors.md">Working with selectors</a> for more details.
		/// If no elements match the selector, returns empty array.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		QuerySelectorAll(*string selector) IReadOnlyCollection<ElementHandle>
		/// <summary>
		/// <para>Returns the buffer with the captured screenshot.</para>
		/// <para>
		/// This method waits for the <a href="./actionability.md">actionability</a> checks,
		/// then scrolls element into view before taking a screenshot. If the element is detached
		/// from DOM, the method throws an error.
		/// </para>
		/// </summary>
		/// <param name="omitBackground">
		/// Hides default white background and allows capturing screenshots with transparency.
		/// Not applicable to <c>jpeg</c> images. Defaults to <c>false</c>.
		/// </param>
		/// <param name="path">
		/// The file path to save the image to. The screenshot type will be inferred from file
		/// extension. If <paramref name="path"/> is a relative path, then it is resolved relative
		/// to the current working directory. If no path is provided, the image won't be saved
		/// to the disk.
		/// </param>
		/// <param name="quality">The quality of the image, between 0-100. Not applicable to <c>png</c> images.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="type">Specify screenshot type, defaults to <c>png</c>.</param>
		Screenshot(*bool omitBackground = default, *string path = default, *int quality = default, *float64 timeout = default, *ScreenshotType type = default) []byte
		/// <summary>
		/// <para>
		/// This method waits for <a href="./actionability.md">actionability</a> checks, then
		/// tries to scroll element into view, unless it is completely visible as defined by
		/// <a href="https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API">IntersectionObserver</a>'s
		/// <c>ratio</c>.
		/// </para>
		/// <para>
		/// Throws when <c>elementHandle</c> does not point to an element <a href="https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected">connected</a>
		/// to a Document or a ShadowRoot.
		/// </para>
		/// </summary>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		ScrollIntoViewIfNeeded(*float64 timeout = default) void
		/// <summary>
		/// <para>Returns the array of option values that have been successfully selected.</para>
		/// <para>
		/// Triggers a <c>change</c> and <c>input</c> event once all the provided options have
		/// been selected. If element is not a <c>&lt;select&gt;</c> element, the method throws
		/// an error.
		/// </para>
		/// <para>Will wait until all specified options are present in the <c>&lt;select&gt;</c> element.</para>
		/// </summary>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SelectOption(*bool noWaitAfter = default, *float64 timeout = default) IReadOnlyCollection<*string>
		/// <summary>
		/// <para>
		/// This method waits for <a href="./actionability.md">actionability</a> checks, then
		/// focuses the element and selects all its text content.
		/// </para>
		/// </summary>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SelectText(*float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method expects <c>elementHandle</c> to point to an <a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input">input
		/// element</a>.
		/// </para>
		/// <para>
		/// Sets the value of the file input to these file paths or files. If some of the <c>filePaths</c>
		/// are relative paths, then they are resolved relative to the the current working directory.
		/// For empty array, clears the selected files.
		/// </para>
		/// </summary>
		/// <param name="files">
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SetInputFiles([]string files, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>This method taps the element by performing the following steps:</para>
		/// <list type="ordinal">
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the element, unless
		/// <paramref name="force"/> option is set.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Touchscreen"/> to tap the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// If the element is detached from the DOM at any moment during the action, this method
		/// rejects.
		/// </para>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>elementHandle.tap()</c> requires that the <c>hasTouch</c> option of the browser
		/// context be set to true.
		/// </para>
		/// </remarks>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Tap(*bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *ElementHandlePosition position = default, *float64 timeout = default) void
		/// <summary><para>Returns the <c>node.textContent</c>.</para></summary>
		GetTextContent() *string
		/// <summary>
		/// <para>
		/// Focuses the element, and then sends a <c>keydown</c>, <c>keypress</c>/<c>input</c>,
		/// and <c>keyup</c> event for each character in the text.
		/// </para>
		/// <para>To press a special key, like <c>Control</c> or <c>ArrowDown</c>, use <see cref="ElementHandle.Press"/>.</para>
		/// <para>An example of typing into a text field and then submitting the form:</para>
		/// </summary>
		/// <param name="text">A text to type into a focused element.</param>
		/// <param name="delay">Time to wait between key presses in milliseconds. Defaults to 0.</param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Type(*string text, *float64 delay = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>This method checks the element by performing the following steps:</para>
		/// <list type="ordinal">
		/// <item><description>
		/// Ensure that element is a checkbox or a radio input. If not, this method rejects.
		/// If the element is already unchecked, this method returns immediately.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the element, unless
		/// <paramref name="force"/> option is set.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>Use <see cref="Page.Mouse"/> to click in the center of the element.</description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// <item><description>Ensure that the element is now unchecked. If not, this method rejects.</description></item>
		/// </list>
		/// <para>
		/// If the element is detached from the DOM at any moment during the action, this method
		/// rejects.
		/// </para>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Uncheck(*bool force = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>Returns when the element satisfies the <paramref name="state"/>.</para>
		/// <para>
		/// Depending on the <paramref name="state"/> parameter, this method waits for one of
		/// the <a href="./actionability.md">actionability</a> checks to pass. This method throws
		/// when the element is detached while waiting, unless waiting for the <c>"hidden"</c>
		/// state.
		/// </para>
		/// <list type="bullet">
		/// <item><description><c>"visible"</c> Wait until the element is <a href="./actionability.md#visible">visible</a>.</description></item>
		/// <item><description>
		/// <c>"hidden"</c> Wait until the element is <a href="./actionability.md#visible">not
		/// visible</a> or <a href="./actionability.md#attached">not attached</a>. Note that
		/// waiting for hidden does not throw when the element detaches.
		/// </description></item>
		/// <item><description>
		/// <c>"stable"</c> Wait until the element is both <a href="./actionability.md#visible">visible</a>
		/// and <a href="./actionability.md#stable">stable</a>.
		/// </description></item>
		/// <item><description><c>"enabled"</c> Wait until the element is <a href="./actionability.md#enabled">enabled</a>.</description></item>
		/// <item><description>
		/// <c>"disabled"</c> Wait until the element is <a href="./actionability.md#enabled">not
		/// enabled</a>.
		/// </description></item>
		/// <item><description><c>"editable"</c> Wait until the element is <a href="./actionability.md#editable">editable</a>.</description></item>
		/// </list>
		/// <para>
		/// If the element does not satisfy the condition for the <paramref name="timeout"/>
		/// milliseconds, this method will throw.
		/// </para>
		/// </summary>
		/// <param name="state">A state to wait for, see below for more details.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForElementState(*ElementState state, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Returns element specified by selector when it satisfies <paramref name="state"/>
		/// option. Returns <c>null</c> if waiting for <c>hidden</c> or <c>detached</c>.
		/// </para>
		/// <para>
		/// Wait for the <paramref name="selector"/> relative to the element handle to satisfy
		/// <paramref name="state"/> option (either appear/disappear from dom, or become visible/hidden).
		/// If at the moment of calling the method <paramref name="selector"/> already satisfies
		/// the condition, the method will return immediately. If the selector doesn't satisfy
		/// the condition for the <paramref name="timeout"/> milliseconds, the function will
		/// throw.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// This method does not work across navigations, use <see cref="Page.WaitForSelector"/>
		/// instead.
		/// </para>
		/// </remarks>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="state">
		/// Defaults to <c>'visible'</c>. Can be either:
		/// <list type="bullet">
		/// <item><description><c>'attached'</c> - wait for element to be present in DOM.</description></item>
		/// <item><description><c>'detached'</c> - wait for element to not be present in DOM.</description></item>
		/// <item><description>
		/// <c>'visible'</c> - wait for element to have non-empty bounding box and no <c>visibility:hidden</c>.
		/// Note that element without any content or with <c>display:none</c> has an empty bounding
		/// box and is not considered visible.
		/// </description></item>
		/// <item><description>
		/// <c>'hidden'</c> - wait for element to be either detached from DOM, or have an empty
		/// bounding box or <c>visibility:hidden</c>. This is opposite to the <c>'visible'</c>
		/// option.
		/// </description></item>
		/// </list>
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForSelector(*string selector, *WaitForSelectorState state = default, *float64 timeout = default) ElementHandle
	}/// <summary>
	/// <para>
	/// <see cref="FileChooser"/> objects are dispatched by the page in the <see cref="Page.FileChooser"/>
	/// event.
	/// </para>
	/// </summary>
	type FileChooser interface {
		/// <summary><para>Returns input element associated with this file chooser.</para></summary>
		ElementHandle Element { get; }
		/// <summary><para>Returns whether this file chooser accepts multiple files.</para></summary>
		*bool IsMultiple { get; }
		/// <summary><para>Returns page this file chooser belongs to.</para></summary>
		Page Page { get; }
		/// <summary>
		/// <para>
		/// Sets the value of the file input this chooser is associated with. If some of the
		/// <c>filePaths</c> are relative paths, then they are resolved relative to the the
		/// current working directory. For empty array, clears the selected files.
		/// </para>
		/// </summary>
		/// <param name="files">
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SetFiles([]string files, *bool noWaitAfter = default, *float64 timeout = default) void
	}/// <summary>
	/// <para>
	/// At every point of time, page exposes its current frame tree via the <see cref="Page.MainFrame"/>
	/// and <see cref="Frame.ChildFrames"/> methods.
	/// </para>
	/// <para>
	/// <see cref="Frame"/> object's lifecycle is controlled by three events, dispatched
	/// on the page object:
	/// </para>
	/// <list type="bullet">
	/// <item><description>
	/// <see cref="Page.FrameAttached"/> - fired when the frame gets attached to the page.
	/// A Frame can be attached to the page only once.
	/// </description></item>
	/// <item><description>
	/// <see cref="Page.FrameNavigated"/> - fired when the frame commits navigation to a
	/// different URL.
	/// </description></item>
	/// <item><description>
	/// <see cref="Page.FrameDetached"/> - fired when the frame gets detached from the page.
	/// A Frame can be detached from the page only once.
	/// </description></item>
	/// </list>
	/// <para>An example of dumping frame tree:</para>
	/// </summary>
	type Frame interface {
		/// <summary>
		/// <para>
		/// Returns the added tag when the script's onload fires or when the script content
		/// was injected into frame.
		/// </para>
		/// <para>Adds a <c>&lt;script&gt;</c> tag into the page with the desired url or content.</para>
		/// </summary>
		/// <param name="content">Raw JavaScript content to be injected into frame.</param>
		/// <param name="path">
		/// Path to the JavaScript file to be injected into frame. If <c>path</c> is a relative
		/// path, then it is resolved relative to the current working directory.
		/// </param>
		/// <param name="type">
		/// Script type. Use 'module' in order to load a Javascript ES6 module. See <a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script">script</a>
		/// for more details.
		/// </param>
		/// <param name="url">URL of a script to be added.</param>
		AddScriptTag(*string content = default, *string path = default, *string type = default, *string url = default) ElementHandle
		/// <summary>
		/// <para>
		/// Returns the added tag when the stylesheet's onload fires or when the CSS content
		/// was injected into frame.
		/// </para>
		/// <para>
		/// Adds a <c>&lt;link rel="stylesheet"&gt;</c> tag into the page with the desired url
		/// or a <c>&lt;style type="text/css"&gt;</c> tag with the content.
		/// </para>
		/// </summary>
		/// <param name="content">Raw CSS content to be injected into frame.</param>
		/// <param name="path">
		/// Path to the CSS file to be injected into frame. If <c>path</c> is a relative path,
		/// then it is resolved relative to the current working directory.
		/// </param>
		/// <param name="url">URL of the <c>&lt;link&gt;</c> tag.</param>
		AddStyleTag(*string content = default, *string path = default, *string url = default) ElementHandle
		/// <summary>
		/// <para>
		/// This method checks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Ensure that matched element is a checkbox or a radio input. If not, this method
		/// rejects. If the element is already checked, this method returns immediately.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>Use <see cref="Page.Mouse"/> to click in the center of the element.</description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// <item><description>Ensure that the element is now checked. If not, this method rejects.</description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Check(*string selector, *bool force = default, *bool noWaitAfter = default, *float64 timeout = default) void
		IReadOnlyCollection<Frame> ChildFrames { get; }
		/// <summary>
		/// <para>
		/// This method clicks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to click in the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="clickCount">defaults to 1. See <see cref="UIEvent.detail"/>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Click(*string selector, *MouseButton button = default, *int clickCount = default, *float64 delay = default, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *FramePosition position = default, *float64 timeout = default) void
		/// <summary><para>Gets the full HTML contents of the frame, including the doctype.</para></summary>
		GetContent() *string
		/// <summary>
		/// <para>
		/// This method double clicks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to double click in the center of the element, or the
		/// specified <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set. Note that if the first click of the <c>dblclick()</c> triggers a
		/// navigation event, this method will reject.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>frame.dblclick()</c> dispatches two <c>click</c> events and a single <c>dblclick</c>
		/// event.
		/// </para>
		/// </remarks>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Dblclick(*string selector, *MouseButton button = default, *float64 delay = default, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *FramePosition position = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// The snippet below dispatches the <c>click</c> event on the element. Regardless of
		/// the visibility state of the elment, <c>click</c> is dispatched. This is equivalend
		/// to calling <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click">element.click()</a>.
		/// </para>
		/// <para>
		/// Under the hood, it creates an instance of an event based on the given <paramref
		/// name="type"/>, initializes it with <paramref name="eventInit"/> properties and dispatches
		/// it on the element. Events are <c>composed</c>, <c>cancelable</c> and bubble by default.
		/// </para>
		/// <para>
		/// Since <paramref name="eventInit"/> is event-specific, please refer to the events
		/// documentation for the lists of initial properties:
		/// </para>
		/// <list type="bullet">
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent">DragEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent">FocusEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent">KeyboardEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent">MouseEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent">PointerEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent">TouchEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/Event/Event">Event</a></description></item>
		/// </list>
		/// <para>
		/// You can also specify <c>JSHandle</c> as the property value if you want live objects
		/// to be passed into the event:
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="type">DOM event type: <c>"click"</c>, <c>"dragstart"</c>, etc.</param>
		/// <param name="eventInit">Optional event-specific initialization properties.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		DispatchEvent(*string selector, *string type, interface{} eventInit = default, *float64 timeout = default) void
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>
		/// The method finds an element matching the specified selector within the frame and
		/// passes it as a first argument to <paramref name="expression"/>. See <a href="./selectors.md">Working
		/// with selectors</a> for more details. If no elements match the selector, the method
		/// throws an error.
		/// </para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <see cref="Frame.EvalOnSelector"/>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvalOnSelector(*string selector, *string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>
		/// The method finds all elements matching the specified selector within the frame and
		/// passes an array of matched elements as a first argument to <paramref name="expression"/>.
		/// See <a href="./selectors.md">Working with selectors</a> for more details.
		/// </para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <see cref="Frame.EvalOnSelectorAll"/>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvalOnSelectorAll(*string selector, *string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>
		/// If the function passed to the <see cref="Frame.Evaluate"/> returns a <see cref="Promise"/>,
		/// then <see cref="Frame.Evaluate"/> would wait for the promise to resolve and return
		/// its value.
		/// </para>
		/// <para>
		/// If the function passed to the <see cref="Frame.Evaluate"/> returns a non-<see cref="Serializable"/>
		/// value, then <see cref="Frame.Evaluate"/> returns <c>undefined</c>. Playwright also
		/// supports transferring some additional values that are not serializable by <c>JSON</c>:
		/// <c>-0</c>, <c>NaN</c>, <c>Infinity</c>, <c>-Infinity</c>.
		/// </para>
		/// <para>A string can also be passed in instead of a function.</para>
		/// <para><see cref="ElementHandle"/> instances can be passed as an argument to the <see cref="Frame.Evaluate"/>:</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		Evaluate(*string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/> as a <see cref="JSHandle"/>.</para>
		/// <para>
		/// The only difference between <see cref="Frame.Evaluate"/> and <see cref="Frame.EvaluateHandle"/>
		/// is that <see cref="Frame.EvaluateHandle"/> returns <see cref="JSHandle"/>.
		/// </para>
		/// <para>
		/// If the function, passed to the <see cref="Frame.EvaluateHandle"/>, returns a <see
		/// cref="Promise"/>, then <see cref="Frame.EvaluateHandle"/> would wait for the promise
		/// to resolve and return its value.
		/// </para>
		/// <para>A string can also be passed in instead of a function.</para>
		/// <para><see cref="JSHandle"/> instances can be passed as an argument to the <see cref="Frame.EvaluateHandle"/>:</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvaluateHandle(*string expression, interface{} arg = default) JSHandle
		/// <summary>
		/// <para>
		/// This method waits for an element matching <paramref name="selector"/>, waits for
		/// <a href="./actionability.md">actionability</a> checks, focuses the element, fills
		/// it and triggers an <c>input</c> event after filling. If the element is inside the
		/// <c>&lt;label&gt;</c> element that has associated <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control">control</a>,
		/// that control will be filled instead. If the element to be filled is not an <c>&lt;input&gt;</c>,
		/// <c>&lt;textarea&gt;</c> or <c>[contenteditable]</c> element, this method throws
		/// an error. Note that you can pass an empty string to clear the input field.
		/// </para>
		/// <para>To send fine-grained keyboard events, use <see cref="Frame.Type"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="value">
		/// Value to fill for the <c>&lt;input&gt;</c>, <c>&lt;textarea&gt;</c> or <c>[contenteditable]</c>
		/// element.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Fill(*string selector, *string value, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method fetches an element with <paramref name="selector"/> and focuses it.
		/// If there's no element matching <paramref name="selector"/>, the method waits until
		/// a matching element appears in the DOM.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Focus(*string selector, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Returns the <c>frame</c> or <c>iframe</c> element handle which corresponds to this
		/// frame.
		/// </para>
		/// <para>
		/// This is an inverse of <see cref="ElementHandle.ContentFrame"/>. Note that returned
		/// handle actually belongs to the parent frame.
		/// </para>
		/// <para>
		/// This method throws an error if the frame has been detached before <c>frameElement()</c>
		/// returns.
		/// </para>
		/// </summary>
		GetFrameElement() ElementHandle
		/// <summary><para>Returns element attribute value.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="name">Attribute name to get the value for.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		GetAttribute(*string selector, *string name, *float64 timeout = default) *string
		/// <summary>
		/// <para>
		/// Returns the main resource response. In case of multiple redirects, the navigation
		/// will resolve with the response of the last redirect.
		/// </para>
		/// <para><c>frame.goto</c> will throw an error if:</para>
		/// <list type="bullet">
		/// <item><description>there's an SSL error (e.g. in case of self-signed certificates).</description></item>
		/// <item><description>target URL is invalid.</description></item>
		/// <item><description>the <paramref name="timeout"/> is exceeded during navigation.</description></item>
		/// <item><description>the remote server does not respond or is unreachable.</description></item>
		/// <item><description>the main resource failed to load.</description></item>
		/// </list>
		/// <para>
		/// <c>frame.goto</c> will not throw an error when any valid HTTP status code is returned
		/// by the remote server, including 404 "Not Found" and 500 "Internal Server Error".
		/// The status code for such responses can be retrieved by calling <see cref="Response.Status"/>.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>frame.goto</c> either throws an error or returns a main resource response. The
		/// only exceptions are navigation to <c>about:blank</c> or navigation to the same URL
		/// with a different hash, which would succeed and return <c>null</c>.
		/// </para>
		/// <para>
		/// Headless mode doesn't support navigation to a PDF document. See the <a href="https://bugs.chromium.org/p/chromium/issues/detail?id=761295">upstream
		/// issue</a>.
		/// </para>
		/// </remarks>
		/// <param name="url">URL to navigate frame to. The url should include scheme, e.g. <c>https://</c>.</param>
		/// <param name="referer">
		/// Referer header value. If provided it will take preference over the referer header
		/// value set by <see cref="Page.SetExtraHttpHeaders"/>.
		/// </param>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		Goto(*string url, *string referer = default, *float64 timeout = default, *WaitUntilState waitUntil = default) Response
		/// <summary>
		/// <para>
		/// This method hovers over an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to hover over the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <c>noWaitAfter</c>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Hover(*string selector, *bool force = default, []KeyboardModifier modifiers = default, *FramePosition position = default, *float64 timeout = default) void
		/// <summary><para>Returns <c>element.innerHTML</c>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		InnerHTML(*string selector, *float64 timeout = default) *string
		/// <summary><para>Returns <c>element.innerText</c>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		InnerText(*string selector, *float64 timeout = default) *string
		/// <summary>
		/// <para>
		/// Returns whether the element is checked. Throws if the element is not a checkbox
		/// or radio input.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsChecked(*string selector, *float64 timeout = default) *bool
		/// <summary><para>Returns <c>true</c> if the frame has been detached, or <c>false</c> otherwise.</para></summary>
		*bool IsDetached { get; }
		/// <summary><para>Returns whether the element is disabled, the opposite of <a href="./actionability.md#enabled">enabled</a>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsDisabled(*string selector, *float64 timeout = default) *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#editable">editable</a>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsEditable(*string selector, *float64 timeout = default) *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#enabled">enabled</a>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsEnabled(*string selector, *float64 timeout = default) *bool
		/// <summary>
		/// <para>
		/// Returns whether the element is hidden, the opposite of <a href="./actionability.md#visible">visible</a>.
		/// <paramref name="selector"/> that does not match any elements is considered hidden.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsHidden(*string selector, *float64 timeout = default) *bool
		/// <summary>
		/// <para>
		/// Returns whether the element is <a href="./actionability.md#visible">visible</a>.
		/// <paramref name="selector"/> that does not match any elements is considered not visible.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsVisible(*string selector, *float64 timeout = default) *bool
		/// <summary>
		/// <para>Returns frame's name attribute as specified in the tag.</para>
		/// <para>If the name is empty, returns the id attribute instead.</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// This value is calculated once when the frame is created, and will not update if
		/// the attribute is changed later.
		/// </para>
		/// </remarks>
		*string Name { get; }
		/// <summary><para>Returns the page containing this frame.</para></summary>
		Page Page { get; }
		/// <summary><para>Parent frame, if any. Detached frames and main frames return <c>null</c>.</para></summary>
		Frame ParentFrame { get; }
		/// <summary>
		/// <para>
		/// <paramref name="key"/> can specify the intended <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key">keyboardEvent.key</a>
		/// value or a single character to generate the text for. A superset of the <paramref
		/// name="key"/> values can be found <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values">here</a>.
		/// Examples of the keys are:
		/// </para>
		/// <para>
		/// <c>F1</c> - <c>F12</c>, <c>Digit0</c>- <c>Digit9</c>, <c>KeyA</c>- <c>KeyZ</c>,
		/// <c>Backquote</c>, <c>Minus</c>, <c>Equal</c>, <c>Backslash</c>, <c>Backspace</c>,
		/// <c>Tab</c>, <c>Delete</c>, <c>Escape</c>, <c>ArrowDown</c>, <c>End</c>, <c>Enter</c>,
		/// <c>Home</c>, <c>Insert</c>, <c>PageDown</c>, <c>PageUp</c>, <c>ArrowRight</c>, <c>ArrowUp</c>,
		/// etc.
		/// </para>
		/// <para>
		/// Following modification shortcuts are also supported: <c>Shift</c>, <c>Control</c>,
		/// <c>Alt</c>, <c>Meta</c>, <c>ShiftLeft</c>.
		/// </para>
		/// <para>
		/// Holding down <c>Shift</c> will type the text that corresponds to the <paramref name="key"/>
		/// in the upper case.
		/// </para>
		/// <para>
		/// If <paramref name="key"/> is a single character, it is case-sensitive, so the values
		/// <c>a</c> and <c>A</c> will generate different respective texts.
		/// </para>
		/// <para>
		/// Shortcuts such as <c>key: "Control+o"</c> or <c>key: "Control+Shift+T"</c> are supported
		/// as well. When speficied with the modifier, modifier is pressed and being held while
		/// the subsequent key is being pressed.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="key">
		/// Name of the key to press or a character to generate, such as <c>ArrowLeft</c> or
		/// <c>a</c>.
		/// </param>
		/// <param name="delay">
		/// Time to wait between <c>keydown</c> and <c>keyup</c> in milliseconds. Defaults to
		/// 0.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Press(*string selector, *string key, *float64 delay = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>Returns the ElementHandle pointing to the frame element.</para>
		/// <para>
		/// The method finds an element matching the specified selector within the frame. See
		/// <a href="./selectors.md">Working with selectors</a> for more details. If no elements
		/// match the selector, returns <c>null</c>.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		QuerySelector(*string selector) ElementHandle
		/// <summary>
		/// <para>Returns the ElementHandles pointing to the frame elements.</para>
		/// <para>
		/// The method finds all elements matching the specified selector within the frame.
		/// See <a href="./selectors.md">Working with selectors</a> for more details. If no
		/// elements match the selector, returns empty array.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		QuerySelectorAll(*string selector) IReadOnlyCollection<ElementHandle>
		/// <summary>
		/// <para>Returns the array of option values that have been successfully selected.</para>
		/// <para>
		/// Triggers a <c>change</c> and <c>input</c> event once all the provided options have
		/// been selected. If there's no <c>&lt;select&gt;</c> element matching <paramref name="selector"/>,
		/// the method throws an error.
		/// </para>
		/// <para>Will wait until all specified options are present in the <c>&lt;select&gt;</c> element.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SelectOption(*string selector, *bool noWaitAfter = default, *float64 timeout = default) IReadOnlyCollection<*string>
		/// <param name="html">HTML markup to assign to the page.</param>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		SetContent(*string html, *float64 timeout = default, *WaitUntilState waitUntil = default) void
		/// <summary>
		/// <para>
		/// This method expects <paramref name="selector"/> to point to an <a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input">input
		/// element</a>.
		/// </para>
		/// <para>
		/// Sets the value of the file input to these file paths or files. If some of the <c>filePaths</c>
		/// are relative paths, then they are resolved relative to the the current working directory.
		/// For empty array, clears the selected files.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="files">
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SetInputFiles(*string selector, []string files, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method taps an element matching <paramref name="selector"/> by performing the
		/// following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Touchscreen"/> to tap the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>frame.tap()</c> requires that the <c>hasTouch</c> option of the browser context
		/// be set to true.
		/// </para>
		/// </remarks>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Tap(*string selector, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *FramePosition position = default, *float64 timeout = default) void
		/// <summary><para>Returns <c>element.textContent</c>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		TextContent(*string selector, *float64 timeout = default) *string
		/// <summary><para>Returns the page title.</para></summary>
		GetTitle() *string
		/// <summary>
		/// <para>
		/// Sends a <c>keydown</c>, <c>keypress</c>/<c>input</c>, and <c>keyup</c> event for
		/// each character in the text. <c>frame.type</c> can be used to send fine-grained keyboard
		/// events. To fill values in form fields, use <see cref="Frame.Fill"/>.
		/// </para>
		/// <para>To press a special key, like <c>Control</c> or <c>ArrowDown</c>, use <see cref="Keyboard.Press"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="text">A text to type into a focused element.</param>
		/// <param name="delay">Time to wait between key presses in milliseconds. Defaults to 0.</param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Type(*string selector, *string text, *float64 delay = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method checks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Ensure that matched element is a checkbox or a radio input. If not, this method
		/// rejects. If the element is already unchecked, this method returns immediately.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>Use <see cref="Page.Mouse"/> to click in the center of the element.</description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// <item><description>Ensure that the element is now unchecked. If not, this method rejects.</description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Uncheck(*string selector, *bool force = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary><para>Returns frame's url.</para></summary>
		*string URL { get; }
		/// <summary>
		/// <para>
		/// Returns when the <paramref name="expression"/> returns a truthy value, returns that
		/// value.
		/// </para>
		/// <para>The <see cref="Frame.WaitForFunction"/> can be used to observe viewport size change:</para>
		/// <para>To pass an argument to the predicate of <c>frame.waitForFunction</c> function:</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		/// <param name="polling">
		/// If <paramref name="polling"/> is <c>'raf'</c>, then <paramref name="expression"/>
		/// is constantly executed in <c>requestAnimationFrame</c> callback. If <paramref name="polling"/>
		/// is a number, then it is treated as an interval in milliseconds at which the function
		/// would be executed. Defaults to <c>raf</c>.
		/// </param>
		/// <param name="timeout">
		/// maximum time to wait for in milliseconds. Defaults to <c>30000</c> (30 seconds).
		/// Pass <c>0</c> to disable timeout. The default value can be changed by using the
		/// <see cref="BrowserContext.SetDefaultTimeout"/>.
		/// </param>
		WaitForFunction(*string expression, interface{} arg = default, interface{} polling = default, *float64 timeout = default) JSHandle
		/// <summary>
		/// <para>Waits for the required load state to be reached.</para>
		/// <para>
		/// This returns when the frame reaches a required load state, <c>load</c> by default.
		/// The navigation must have been committed when this method is called. If current document
		/// has already reached the required state, resolves immediately.
		/// </para>
		/// </summary>
		/// <param name="state">
		/// Optional load state to wait for, defaults to <c>load</c>. If the state has been
		/// already reached while loading current document, the method resolves immediately.
		/// Can be one of:
		/// <list type="bullet">
		/// <item><description><c>'load'</c> - wait for the <c>load</c> event to be fired.</description></item>
		/// <item><description><c>'domcontentloaded'</c> - wait for the <c>DOMContentLoaded</c> event to be fired.</description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - wait until there are no network connections for at least
		/// <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForLoadState(*LoadState state = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Waits for the frame navigation and returns the main resource response. In case of
		/// multiple redirects, the navigation will resolve with the response of the last redirect.
		/// In case of navigation to a different anchor or navigation due to History API usage,
		/// the navigation will resolve with <c>null</c>.
		/// </para>
		/// <para>
		/// This method waits for the frame to navigate to a new URL. It is useful for when
		/// you run code which will indirectly cause the frame to navigate. Consider this example:
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Usage of the <a href="https://developer.mozilla.org/en-US/docs/Web/API/History_API">History
		/// API</a> to change the URL is considered a navigation.
		/// </para>
		/// </remarks>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="url">
		/// A glob pattern, regex pattern or predicate receiving <see cref="URL"/> to match
		/// while waiting for the navigation.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		WaitForNavigation(*float64 timeout = default, interface{} url = default, *WaitUntilState waitUntil = default) Response
		/// <summary>
		/// <para>
		/// Returns when element specified by selector satisfies <paramref name="state"/> option.
		/// Returns <c>null</c> if waiting for <c>hidden</c> or <c>detached</c>.
		/// </para>
		/// <para>
		/// Wait for the <paramref name="selector"/> to satisfy <paramref name="state"/> option
		/// (either appear/disappear from dom, or become visible/hidden). If at the moment of
		/// calling the method <paramref name="selector"/> already satisfies the condition,
		/// the method will return immediately. If the selector doesn't satisfy the condition
		/// for the <paramref name="timeout"/> milliseconds, the function will throw.
		/// </para>
		/// <para>This method works across navigations:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="state">
		/// Defaults to <c>'visible'</c>. Can be either:
		/// <list type="bullet">
		/// <item><description><c>'attached'</c> - wait for element to be present in DOM.</description></item>
		/// <item><description><c>'detached'</c> - wait for element to not be present in DOM.</description></item>
		/// <item><description>
		/// <c>'visible'</c> - wait for element to have non-empty bounding box and no <c>visibility:hidden</c>.
		/// Note that element without any content or with <c>display:none</c> has an empty bounding
		/// box and is not considered visible.
		/// </description></item>
		/// <item><description>
		/// <c>'hidden'</c> - wait for element to be either detached from DOM, or have an empty
		/// bounding box or <c>visibility:hidden</c>. This is opposite to the <c>'visible'</c>
		/// option.
		/// </description></item>
		/// </list>
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForSelector(*string selector, *WaitForSelectorState state = default, *float64 timeout = default) ElementHandle
		/// <summary>
		/// <para>Waits for the given <paramref name="timeout"/> in milliseconds.</para>
		/// <para>
		/// Note that <c>frame.waitForTimeout()</c> should only be used for debugging. Tests
		/// using the timer in production are going to be flaky. Use signals such as network
		/// events, selectors becoming visible and others instead.
		/// </para>
		/// </summary>
		/// <param name="timeout">A timeout to wait for</param>
		WaitForTimeout(*float64 timeout) void
	}/// <summary>
	/// <para>
	/// JSHandle represents an in-page JavaScript object. JSHandles can be created with
	/// the <see cref="Page.EvaluateHandle"/> method.
	/// </para>
	/// <para>
	/// JSHandle prevents the referenced JavaScript object being garbage collected unless
	/// the handle is exposed with <see cref="JSHandle.Dispose"/>. JSHandles are auto-disposed
	/// when their origin frame gets navigated or the parent context gets destroyed.
	/// </para>
	/// <para>
	/// JSHandle instances can be used as an argument in <see cref="Page.EvalOnSelector"/>,
	/// <see cref="Page.Evaluate"/> and <see cref="Page.EvaluateHandle"/> methods.
	/// </para>
	/// </summary>
	type JSHandle interface {
		/// <summary>
		/// <para>
		/// Returns either <c>null</c> or the object handle itself, if the object handle is
		/// an instance of <see cref="ElementHandle"/>.
		/// </para>
		/// </summary>
		ElementHandle AsElement { get; }
		/// <summary><para>The <c>jsHandle.dispose</c> method stops referencing the element handle.</para></summary>
		Dispose() void
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>This method passes this handle as the first argument to <paramref name="expression"/>.</para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <c>handle.evaluate</c>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		Evaluate(*string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/> as a <see cref="JSHandle"/>.</para>
		/// <para>This method passes this handle as the first argument to <paramref name="expression"/>.</para>
		/// <para>
		/// The only difference between <c>jsHandle.evaluate</c> and <c>jsHandle.evaluateHandle</c>
		/// is that <c>jsHandle.evaluateHandle</c> returns <see cref="JSHandle"/>.
		/// </para>
		/// <para>
		/// If the function passed to the <c>jsHandle.evaluateHandle</c> returns a <see cref="Promise"/>,
		/// then <c>jsHandle.evaluateHandle</c> would wait for the promise to resolve and return
		/// its value.
		/// </para>
		/// <para>See <see cref="Page.EvaluateHandle"/> for more details.</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvaluateHandle(*string expression, interface{} arg = default) JSHandle
		/// <summary>
		/// <para>
		/// The method returns a map with **own property names** as keys and JSHandle instances
		/// for the property values.
		/// </para>
		/// </summary>
		GetProperties() Dictionary<*string, JSHandle>
		/// <summary><para>Fetches a single property from the referenced object.</para></summary>
		/// <param name="propertyName">property to get</param>
		GetProperty(*string propertyName) JSHandle
		/// <summary>
		/// <para>
		/// Returns a JSON representation of the object. If the object has a <c>toJSON</c> function,
		/// it **will not be called**.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// The method will return an empty JSON object if the referenced object is not stringifiable.
		/// It will throw an error if the object has circular references.
		/// </para>
		/// </remarks>
		GetJsonValue() interface{}
	}/// <summary>
	/// <para>
	/// Keyboard provides an api for managing a virtual keyboard. The high level api is
	/// <see cref="Keyboard.Type"/>, which takes raw characters and generates proper keydown,
	/// keypress/input, and keyup events on your page.
	/// </para>
	/// <para>
	/// For finer control, you can use <see cref="Keyboard.Down"/>, <see cref="Keyboard.Up"/>,
	/// and <see cref="Keyboard.InsertText"/> to manually fire events as if they were generated
	/// from a real keyboard.
	/// </para>
	/// <para>An example of holding down <c>Shift</c> in order to select and delete some text:</para>
	/// <para>An example of pressing uppercase <c>A</c></para>
	/// <para>An example to trigger select-all with the keyboard</para>
	/// </summary>
	type Keyboard interface {
		/// <summary>
		/// <para>Dispatches a <c>keydown</c> event.</para>
		/// <para>
		/// <paramref name="key"/> can specify the intended <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key">keyboardEvent.key</a>
		/// value or a single character to generate the text for. A superset of the <paramref
		/// name="key"/> values can be found <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values">here</a>.
		/// Examples of the keys are:
		/// </para>
		/// <para>
		/// <c>F1</c> - <c>F12</c>, <c>Digit0</c>- <c>Digit9</c>, <c>KeyA</c>- <c>KeyZ</c>,
		/// <c>Backquote</c>, <c>Minus</c>, <c>Equal</c>, <c>Backslash</c>, <c>Backspace</c>,
		/// <c>Tab</c>, <c>Delete</c>, <c>Escape</c>, <c>ArrowDown</c>, <c>End</c>, <c>Enter</c>,
		/// <c>Home</c>, <c>Insert</c>, <c>PageDown</c>, <c>PageUp</c>, <c>ArrowRight</c>, <c>ArrowUp</c>,
		/// etc.
		/// </para>
		/// <para>
		/// Following modification shortcuts are also supported: <c>Shift</c>, <c>Control</c>,
		/// <c>Alt</c>, <c>Meta</c>, <c>ShiftLeft</c>.
		/// </para>
		/// <para>
		/// Holding down <c>Shift</c> will type the text that corresponds to the <paramref name="key"/>
		/// in the upper case.
		/// </para>
		/// <para>
		/// If <paramref name="key"/> is a single character, it is case-sensitive, so the values
		/// <c>a</c> and <c>A</c> will generate different respective texts.
		/// </para>
		/// <para>
		/// If <paramref name="key"/> is a modifier key, <c>Shift</c>, <c>Meta</c>, <c>Control</c>,
		/// or <c>Alt</c>, subsequent key presses will be sent with that modifier active. To
		/// release the modifier key, use <see cref="Keyboard.Up"/>.
		/// </para>
		/// <para>
		/// After the key is pressed once, subsequent calls to <see cref="Keyboard.Down"/> will
		/// have <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/repeat">repeat</a>
		/// set to true. To release the key, use <see cref="Keyboard.Up"/>.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Modifier keys DO influence <c>keyboard.down</c>. Holding down <c>Shift</c> will
		/// type the text in upper case.
		/// </para>
		/// </remarks>
		/// <param name="key">
		/// Name of the key to press or a character to generate, such as <c>ArrowLeft</c> or
		/// <c>a</c>.
		/// </param>
		Down(*string key) void
		/// <summary>
		/// <para>
		/// Dispatches only <c>input</c> event, does not emit the <c>keydown</c>, <c>keyup</c>
		/// or <c>keypress</c> events.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Modifier keys DO NOT effect <c>keyboard.insertText</c>. Holding down <c>Shift</c>
		/// will not type the text in upper case.
		/// </para>
		/// </remarks>
		/// <param name="text">Sets input to the specified text value.</param>
		InsertText(*string text) void
		/// <summary>
		/// <para>
		/// <paramref name="key"/> can specify the intended <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key">keyboardEvent.key</a>
		/// value or a single character to generate the text for. A superset of the <paramref
		/// name="key"/> values can be found <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values">here</a>.
		/// Examples of the keys are:
		/// </para>
		/// <para>
		/// <c>F1</c> - <c>F12</c>, <c>Digit0</c>- <c>Digit9</c>, <c>KeyA</c>- <c>KeyZ</c>,
		/// <c>Backquote</c>, <c>Minus</c>, <c>Equal</c>, <c>Backslash</c>, <c>Backspace</c>,
		/// <c>Tab</c>, <c>Delete</c>, <c>Escape</c>, <c>ArrowDown</c>, <c>End</c>, <c>Enter</c>,
		/// <c>Home</c>, <c>Insert</c>, <c>PageDown</c>, <c>PageUp</c>, <c>ArrowRight</c>, <c>ArrowUp</c>,
		/// etc.
		/// </para>
		/// <para>
		/// Following modification shortcuts are also supported: <c>Shift</c>, <c>Control</c>,
		/// <c>Alt</c>, <c>Meta</c>, <c>ShiftLeft</c>.
		/// </para>
		/// <para>
		/// Holding down <c>Shift</c> will type the text that corresponds to the <paramref name="key"/>
		/// in the upper case.
		/// </para>
		/// <para>
		/// If <paramref name="key"/> is a single character, it is case-sensitive, so the values
		/// <c>a</c> and <c>A</c> will generate different respective texts.
		/// </para>
		/// <para>
		/// Shortcuts such as <c>key: "Control+o"</c> or <c>key: "Control+Shift+T"</c> are supported
		/// as well. When speficied with the modifier, modifier is pressed and being held while
		/// the subsequent key is being pressed.
		/// </para>
		/// <para>Shortcut for <see cref="Keyboard.Down"/> and <see cref="Keyboard.Up"/>.</para>
		/// </summary>
		/// <param name="key">
		/// Name of the key to press or a character to generate, such as <c>ArrowLeft</c> or
		/// <c>a</c>.
		/// </param>
		/// <param name="delay">
		/// Time to wait between <c>keydown</c> and <c>keyup</c> in milliseconds. Defaults to
		/// 0.
		/// </param>
		Press(*string key, *float64 delay = default) void
		/// <summary>
		/// <para>
		/// Sends a <c>keydown</c>, <c>keypress</c>/<c>input</c>, and <c>keyup</c> event for
		/// each character in the text.
		/// </para>
		/// <para>To press a special key, like <c>Control</c> or <c>ArrowDown</c>, use <see cref="Keyboard.Press"/>.</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Modifier keys DO NOT effect <c>keyboard.type</c>. Holding down <c>Shift</c> will
		/// not type the text in upper case.
		/// </para>
		/// </remarks>
		/// <param name="text">A text to type into a focused element.</param>
		/// <param name="delay">Time to wait between key presses in milliseconds. Defaults to 0.</param>
		Type(*string text, *float64 delay = default) void
		/// <summary><para>Dispatches a <c>keyup</c> event.</para></summary>
		/// <param name="key">
		/// Name of the key to press or a character to generate, such as <c>ArrowLeft</c> or
		/// <c>a</c>.
		/// </param>
		Up(*string key) void
	}/// <summary>
	/// <para>
	/// The Mouse class operates in main-frame CSS pixels relative to the top-left corner
	/// of the viewport.
	/// </para>
	/// <para>Every <c>page</c> object has its own Mouse, accessible with <see cref="Page.Mouse"/>.</para>
	/// <code>
	/// await Page.Mouse.MoveAsync(0, 0);<br/>
	/// await Page.Mouse.DownAsync();<br/>
	/// await Page.Mouse.MoveAsync(0, 100);<br/>
	/// await Page.Mouse.MoveAsync(100, 100);<br/>
	/// await Page.Mouse.MoveAsync(100, 0);<br/>
	/// await Page.Mouse.MoveAsync(0, 0);<br/>
	/// await Page.Mouse.UpAsync();
	/// </code>
	/// </summary>
	type Mouse interface {
		/// <summary><para>Shortcut for <see cref="Mouse.Move"/>, <see cref="Mouse.Down"/>, <see cref="Mouse.Up"/>.</para></summary>
		/// <param name="x">
		/// </param>
		/// <param name="y">
		/// </param>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="clickCount">defaults to 1. See <see cref="UIEvent.detail"/>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		Click(*float64 x, *float64 y, *MouseButton button = default, *int clickCount = default, *float64 delay = default) void
		/// <summary>
		/// <para>
		/// Shortcut for <see cref="Mouse.Move"/>, <see cref="Mouse.Down"/>, <see cref="Mouse.Up"/>,
		/// <see cref="Mouse.Down"/> and <see cref="Mouse.Up"/>.
		/// </para>
		/// </summary>
		/// <param name="x">
		/// </param>
		/// <param name="y">
		/// </param>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		Dblclick(*float64 x, *float64 y, *MouseButton button = default, *float64 delay = default) void
		/// <summary><para>Dispatches a <c>mousedown</c> event.</para></summary>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="clickCount">defaults to 1. See <see cref="UIEvent.detail"/>.</param>
		Down(*MouseButton button = default, *int clickCount = default) void
		/// <summary><para>Dispatches a <c>mousemove</c> event.</para></summary>
		/// <param name="x">
		/// </param>
		/// <param name="y">
		/// </param>
		/// <param name="steps">defaults to 1. Sends intermediate <c>mousemove</c> events.</param>
		Move(*float64 x, *float64 y, *int steps = default) void
		/// <summary><para>Dispatches a <c>mouseup</c> event.</para></summary>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="clickCount">defaults to 1. See <see cref="UIEvent.detail"/>.</param>
		Up(*MouseButton button = default, *int clickCount = default) void
	}/// <summary>
	/// <para>
	/// Page provides methods to interact with a single tab in a <see cref="Browser"/>,
	/// or an <a href="https://developer.chrome.com/extensions/background_pages">extension
	/// background page</a> in Chromium. One <see cref="Browser"/> instance might have multiple
	/// <see cref="Page"/> instances.
	/// </para>
	/// <para>This example creates a page, navigates it to a URL, and then saves a screenshot:</para>
	/// <para>
	/// The Page class emits various events (described below) which can be handled using
	/// any of Node's native <a href="https://nodejs.org/api/events.html#events_class_eventemitter"><c>EventEmitter</c></a>
	/// methods, such as <c>on</c>, <c>once</c> or <c>removeListener</c>.
	/// </para>
	/// <para>This example logs a message for a single page <c>load</c> event:</para>
	/// <para>To unsubscribe from events use the <c>removeListener</c> method:</para>
	/// </summary>
	type Page interface {
	EventEmitter
		/// <summary><para>Emitted when the page closes.</para></summary>
		event EventHandler Close;
		/// <summary>
		/// <para>
		/// Emitted when JavaScript within the page calls one of console API methods, e.g. <c>console.log</c>
		/// or <c>console.dir</c>. Also emitted if the page throws an error or a warning.
		/// </para>
		/// <para>The arguments passed into <c>console.log</c> appear as arguments on the event handler.</para>
		/// <para>An example of handling <c>console</c> event:</para>
		/// </summary>
		event EventHandler<ConsoleMessage> Console;
		/// <summary>
		/// <para>
		/// Emitted when the page crashes. Browser pages might crash if they try to allocate
		/// too much memory. When the page crashes, ongoing and subsequent operations will throw.
		/// </para>
		/// <para>The most common way to deal with crashes is to catch an exception:</para>
		/// </summary>
		event EventHandler Crash;
		/// <summary>
		/// <para>
		/// Emitted when a JavaScript dialog appears, such as <c>alert</c>, <c>prompt</c>, <c>confirm</c>
		/// or <c>beforeunload</c>. Listener **must** either <see cref="Dialog.Accept"/> or
		/// <see cref="Dialog.Dismiss"/> the dialog - otherwise the page will <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking">freeze</a>
		/// waiting for the dialog, and actions like click will never finish.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// When no <see cref="Page.Dialog"/> listeners are present, all dialogs are automatically
		/// dismissed.
		/// </para>
		/// </remarks>
		event EventHandler<Dialog> Dialog;
		/// <summary>
		/// <para>
		/// Emitted when the JavaScript <a href="https://developer.mozilla.org/en-US/docs/Web/Events/DOMContentLoaded"><c>DOMContentLoaded</c></a>
		/// event is dispatched.
		/// </para>
		/// </summary>
		event EventHandler DOMContentLoaded;
		/// <summary>
		/// <para>
		/// Emitted when attachment download started. User can access basic file operations
		/// on downloaded content via the passed <see cref="Download"/> instance.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Browser context **must** be created with the <paramref name="acceptDownloads"/>
		/// set to <c>true</c> when user needs access to the downloaded content. If <paramref
		/// name="acceptDownloads"/> is not set, download events are emitted, but the actual
		/// download is not performed and user has no access to the downloaded files.
		/// </para>
		/// </remarks>
		event EventHandler<Download> Download;
		/// <summary>
		/// <para>
		/// Emitted when a file chooser is supposed to appear, such as after clicking the  <c>&lt;input
		/// type=file&gt;</c>. Playwright can respond to it via setting the input files using
		/// <see cref="FileChooser.SetFiles"/> that can be uploaded after that.
		/// </para>
		/// </summary>
		event EventHandler<FileChooser> FileChooser;
		/// <summary><para>Emitted when a frame is attached.</para></summary>
		event EventHandler<Frame> FrameAttached;
		/// <summary><para>Emitted when a frame is detached.</para></summary>
		event EventHandler<Frame> FrameDetached;
		/// <summary><para>Emitted when a frame is navigated to a new url.</para></summary>
		event EventHandler<Frame> FrameNavigated;
		/// <summary>
		/// <para>
		/// Emitted when the JavaScript <a href="https://developer.mozilla.org/en-US/docs/Web/Events/load"><c>load</c></a>
		/// event is dispatched.
		/// </para>
		/// </summary>
		event EventHandler Load;
		/// <summary><para>Emitted when an uncaught exception happens within the page.</para></summary>
		event EventHandler<Exception> PageError;
		/// <summary>
		/// <para>
		/// Emitted when the page opens a new tab or window. This event is emitted in addition
		/// to the <see cref="BrowserContext.Page"/>, but only for popups relevant to this page.
		/// </para>
		/// <para>
		/// The earliest moment that page is available is when it has navigated to the initial
		/// url. For example, when opening a popup with <c>window.open('http://example.com')</c>,
		/// this event will fire when the network request to "http://example.com" is done and
		/// its response has started loading in the popup.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Use <see cref="Page.WaitForLoadState"/> to wait until the page gets to a particular
		/// state (you should not need it in most cases).
		/// </para>
		/// </remarks>
		event EventHandler Popup;
		/// <summary>
		/// <para>
		/// Emitted when a page issues a request. The <see cref="request"/> object is read-only.
		/// In order to intercept and mutate requests, see <see cref="Page.Route"/> or <see
		/// cref="BrowserContext.Route"/>.
		/// </para>
		/// </summary>
		event EventHandler<Request> Request;
		/// <summary><para>Emitted when a request fails, for example by timing out.</para></summary>
		/// <remarks>
		/// <para>
		/// HTTP Error responses, such as 404 or 503, are still successful responses from HTTP
		/// standpoint, so request will complete with <see cref="Page.RequestFinished"/> event
		/// and not with <see cref="Page.RequestFailed"/>.
		/// </para>
		/// </remarks>
		event EventHandler<Request> RequestFailed;
		/// <summary>
		/// <para>
		/// Emitted when a request finishes successfully after downloading the response body.
		/// For a successful response, the sequence of events is <c>request</c>, <c>response</c>
		/// and <c>requestfinished</c>.
		/// </para>
		/// </summary>
		event EventHandler<Request> RequestFinished;
		/// <summary>
		/// <para>
		/// Emitted when <see cref="response"/> status and headers are received for a request.
		/// For a successful response, the sequence of events is <c>request</c>, <c>response</c>
		/// and <c>requestfinished</c>.
		/// </para>
		/// </summary>
		event EventHandler<Response> Response;
		/// <summary><para>Emitted when <see cref="WebSocket"/> request is sent.</para></summary>
		event EventHandler<WebSocket> WebSocket;
		/// <summary>
		/// <para>
		/// Emitted when a dedicated <a href="https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API">WebWorker</a>
		/// is spawned by the page.
		/// </para>
		/// </summary>
		event EventHandler<Worker> Worker;
		/// <summary>
		/// <para>Adds a script which would be evaluated in one of the following scenarios:</para>
		/// <list type="bullet">
		/// <item><description>Whenever the page is navigated.</description></item>
		/// <item><description>
		/// Whenever the child frame is attached or navigated. In this case, the script is evaluated
		/// in the context of the newly attached frame.
		/// </description></item>
		/// </list>
		/// <para>
		/// The script is evaluated after the document was created but before any of its scripts
		/// were run. This is useful to amend the JavaScript environment, e.g. to seed <c>Math.random</c>.
		/// </para>
		/// <para>An example of overriding <c>Math.random</c> before the page loads:</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// The order of evaluation of multiple scripts installed via <see cref="BrowserContext.AddInitScript"/>
		/// and <see cref="Page.AddInitScript"/> is not defined.
		/// </para>
		/// </remarks>
		/// <param name="script">Optional Script source to be evaluated in all pages in the browser context.</param>
		/// <param name="path">Optional Script path to be evaluated in all pages in the browser context.</param>
		AddInitScript(*string script = default, *string path = default) void
		/// <summary>
		/// <para>
		/// Adds a <c>&lt;script&gt;</c> tag into the page with the desired url or content.
		/// Returns the added tag when the script's onload fires or when the script content
		/// was injected into frame.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.AddScriptTag"/>.</para>
		/// </summary>
		/// <param name="content">Raw JavaScript content to be injected into frame.</param>
		/// <param name="path">
		/// Path to the JavaScript file to be injected into frame. If <c>path</c> is a relative
		/// path, then it is resolved relative to the current working directory.
		/// </param>
		/// <param name="type">
		/// Script type. Use 'module' in order to load a Javascript ES6 module. See <a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script">script</a>
		/// for more details.
		/// </param>
		/// <param name="url">URL of a script to be added.</param>
		AddScriptTag(*string content = default, *string path = default, *string type = default, *string url = default) ElementHandle
		/// <summary>
		/// <para>
		/// Adds a <c>&lt;link rel="stylesheet"&gt;</c> tag into the page with the desired url
		/// or a <c>&lt;style type="text/css"&gt;</c> tag with the content. Returns the added
		/// tag when the stylesheet's onload fires or when the CSS content was injected into
		/// frame.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.AddStyleTag"/>.</para>
		/// </summary>
		/// <param name="content">Raw CSS content to be injected into frame.</param>
		/// <param name="path">
		/// Path to the CSS file to be injected into frame. If <c>path</c> is a relative path,
		/// then it is resolved relative to the current working directory.
		/// </param>
		/// <param name="url">URL of the <c>&lt;link&gt;</c> tag.</param>
		AddStyleTag(*string content = default, *string path = default, *string url = default) ElementHandle
		/// <summary><para>Brings page to front (activates tab).</para></summary>
		BringToFront() void
		/// <summary>
		/// <para>
		/// This method checks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Ensure that matched element is a checkbox or a radio input. If not, this method
		/// rejects. If the element is already checked, this method returns immediately.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>Use <see cref="Page.Mouse"/> to click in the center of the element.</description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// <item><description>Ensure that the element is now checked. If not, this method rejects.</description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Check"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Check(*string selector, *bool force = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method clicks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to click in the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Click"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="clickCount">defaults to 1. See <see cref="UIEvent.detail"/>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Click(*string selector, *MouseButton button = default, *int clickCount = default, *float64 delay = default, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *PagePosition position = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// If <paramref name="runBeforeUnload"/> is <c>false</c>, does not run any unload handlers
		/// and waits for the page to be closed. If <paramref name="runBeforeUnload"/> is <c>true</c>
		/// the method will run unload handlers, but will **not** wait for the page to close.
		/// </para>
		/// <para>By default, <c>page.close()</c> **does not** run <c>beforeunload</c> handlers.</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// if <paramref name="runBeforeUnload"/> is passed as true, a <c>beforeunload</c> dialog
		/// might be summoned and should be handled manually via <see cref="Page.Dialog"/> event.
		/// </para>
		/// </remarks>
		/// <param name="runBeforeUnload">
		/// Defaults to <c>false</c>. Whether to run the <a href="https://developer.mozilla.org/en-US/docs/Web/Events/beforeunload">before
		/// unload</a> page handlers.
		/// </param>
		Close(*bool runBeforeUnload = default) void
		/// <summary><para>Gets the full HTML contents of the page, including the doctype.</para></summary>
		GetContent() *string
		/// <summary><para>Get the browser context that the page belongs to.</para></summary>
		BrowserContext Context { get; }
		/// <summary>
		/// <para>
		/// This method double clicks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to double click in the center of the element, or the
		/// specified <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set. Note that if the first click of the <c>dblclick()</c> triggers a
		/// navigation event, this method will reject.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Dblclick"/>.</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>page.dblclick()</c> dispatches two <c>click</c> events and a single <c>dblclick</c>
		/// event.
		/// </para>
		/// </remarks>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="button">Defaults to <c>left</c>.</param>
		/// <param name="delay">
		/// Time to wait between <c>mousedown</c> and <c>mouseup</c> in milliseconds. Defaults
		/// to 0.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Dblclick(*string selector, *MouseButton button = default, *float64 delay = default, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *PagePosition position = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// The snippet below dispatches the <c>click</c> event on the element. Regardless of
		/// the visibility state of the elment, <c>click</c> is dispatched. This is equivalend
		/// to calling <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click">element.click()</a>.
		/// </para>
		/// <para>
		/// Under the hood, it creates an instance of an event based on the given <paramref
		/// name="type"/>, initializes it with <paramref name="eventInit"/> properties and dispatches
		/// it on the element. Events are <c>composed</c>, <c>cancelable</c> and bubble by default.
		/// </para>
		/// <para>
		/// Since <paramref name="eventInit"/> is event-specific, please refer to the events
		/// documentation for the lists of initial properties:
		/// </para>
		/// <list type="bullet">
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent">DragEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent">FocusEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent">KeyboardEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent">MouseEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent">PointerEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent">TouchEvent</a></description></item>
		/// <item><description><a href="https://developer.mozilla.org/en-US/docs/Web/API/Event/Event">Event</a></description></item>
		/// </list>
		/// <para>
		/// You can also specify <c>JSHandle</c> as the property value if you want live objects
		/// to be passed into the event:
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="type">DOM event type: <c>"click"</c>, <c>"dragstart"</c>, etc.</param>
		/// <param name="eventInit">Optional event-specific initialization properties.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		DispatchEvent(*string selector, *string type, interface{} eventInit = default, *float64 timeout = default) void
		/// <param name="colorScheme">
		/// Emulates <c>'prefers-colors-scheme'</c> media feature, supported values are <c>'light'</c>,
		/// <c>'dark'</c>, <c>'no-preference'</c>. Passing <c>null</c> disables color scheme
		/// emulation.
		/// </param>
		/// <param name="media">
		/// Changes the CSS media type of the page. The only allowed values are <c>'screen'</c>,
		/// <c>'print'</c> and <c>null</c>. Passing <c>null</c> disables CSS media emulation.
		/// </param>
		EmulateMedia(*ColorScheme colorScheme = default, *Media media = default) void
		/// <summary>
		/// <para>
		/// The method finds an element matching the specified selector within the page and
		/// passes it as a first argument to <paramref name="expression"/>. If no elements match
		/// the selector, the method throws an error. Returns the value of <paramref name="expression"/>.
		/// </para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <see cref="Page.EvalOnSelector"/>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// <para>Shortcut for main frame's <see cref="Frame.EvalOnSelector"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvalOnSelector(*string selector, *string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>
		/// The method finds all elements matching the specified selector within the page and
		/// passes an array of matched elements as a first argument to <paramref name="expression"/>.
		/// Returns the result of <paramref name="expression"/> invocation.
		/// </para>
		/// <para>
		/// If <paramref name="expression"/> returns a <see cref="Promise"/>, then <see cref="Page.EvalOnSelectorAll"/>
		/// would wait for the promise to resolve and return its value.
		/// </para>
		/// <para>Examples:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvalOnSelectorAll(*string selector, *string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the value of the <paramref name="expression"/> invocation.</para>
		/// <para>
		/// If the function passed to the <see cref="Page.Evaluate"/> returns a <see cref="Promise"/>,
		/// then <see cref="Page.Evaluate"/> would wait for the promise to resolve and return
		/// its value.
		/// </para>
		/// <para>
		/// If the function passed to the <see cref="Page.Evaluate"/> returns a non-<see cref="Serializable"/>
		/// value, then <see cref="Page.Evaluate"/> resolves to <c>undefined</c>. Playwright
		/// also supports transferring some additional values that are not serializable by <c>JSON</c>:
		/// <c>-0</c>, <c>NaN</c>, <c>Infinity</c>, <c>-Infinity</c>.
		/// </para>
		/// <para>Passing argument to <paramref name="expression"/>:</para>
		/// <para>A string can also be passed in instead of a function:</para>
		/// <para><see cref="ElementHandle"/> instances can be passed as an argument to the <see cref="Page.Evaluate"/>:</para>
		/// <para>Shortcut for main frame's <see cref="Frame.Evaluate"/>.</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		Evaluate(*string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the value of the <paramref name="expression"/> invocation as a <see cref="JSHandle"/>.</para>
		/// <para>
		/// The only difference between <see cref="Page.Evaluate"/> and <see cref="Page.EvaluateHandle"/>
		/// is that <see cref="Page.EvaluateHandle"/> returns <see cref="JSHandle"/>.
		/// </para>
		/// <para>
		/// If the function passed to the <see cref="Page.EvaluateHandle"/> returns a <see cref="Promise"/>,
		/// then <see cref="Page.EvaluateHandle"/> would wait for the promise to resolve and
		/// return its value.
		/// </para>
		/// <para>A string can also be passed in instead of a function:</para>
		/// <para><see cref="JSHandle"/> instances can be passed as an argument to the <see cref="Page.EvaluateHandle"/>:</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvaluateHandle(*string expression, interface{} arg = default) JSHandle
		/// <summary>
		/// <para>
		/// The method adds a function called <paramref name="name"/> on the <c>window</c> object
		/// of every frame in this page. When called, the function executes <paramref name="callback"/>
		/// and returns a <see cref="Promise"/> which resolves to the return value of <paramref
		/// name="callback"/>. If the <paramref name="callback"/> returns a <see cref="Promise"/>,
		/// it will be awaited.
		/// </para>
		/// <para>
		/// The first argument of the <paramref name="callback"/> function contains information
		/// about the caller: <c>{ browserContext: BrowserContext, page: Page, frame: Frame
		/// }</c>.
		/// </para>
		/// <para>See <see cref="BrowserContext.ExposeBinding"/> for the context-wide version.</para>
		/// <para>An example of exposing page URL to all frames in a page:</para>
		/// <para>An example of passing an element handle:</para>
		/// </summary>
		/// <remarks><para>Functions installed via <see cref="Page.ExposeBinding"/> survive navigations.</para></remarks>
		/// <param name="name">Name of the function on the window object.</param>
		/// <param name="callback">Callback function that will be called in the Playwright's context.</param>
		/// <param name="handle">
		/// Whether to pass the argument as a handle, instead of passing by value. When passing
		/// a handle, only one argument is supported. When passing by value, multiple arguments
		/// are supported.
		/// </param>
		ExposeBinding(*string name, interface{} callback, *bool handle = default) void
		/// <summary>
		/// <para>
		/// The method adds a function called <paramref name="name"/> on the <c>window</c> object
		/// of every frame in the page. When called, the function executes <paramref name="callback"/>
		/// and returns a <see cref="Promise"/> which resolves to the return value of <paramref
		/// name="callback"/>.
		/// </para>
		/// <para>If the <paramref name="callback"/> returns a <see cref="Promise"/>, it will be awaited.</para>
		/// <para>See <see cref="BrowserContext.ExposeFunction"/> for context-wide exposed function.</para>
		/// <para>An example of adding an <c>sha1</c> function to the page:</para>
		/// </summary>
		/// <remarks><para>Functions installed via <see cref="Page.ExposeFunction"/> survive navigations.</para></remarks>
		/// <param name="name">Name of the function on the window object</param>
		/// <param name="callback">Callback function which will be called in Playwright's context.</param>
		ExposeFunction(*string name, interface{} callback) void
		/// <summary>
		/// <para>
		/// This method waits for an element matching <paramref name="selector"/>, waits for
		/// <a href="./actionability.md">actionability</a> checks, focuses the element, fills
		/// it and triggers an <c>input</c> event after filling. If the element is inside the
		/// <c>&lt;label&gt;</c> element that has associated <a href="https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control">control</a>,
		/// that control will be filled instead. If the element to be filled is not an <c>&lt;input&gt;</c>,
		/// <c>&lt;textarea&gt;</c> or <c>[contenteditable]</c> element, this method throws
		/// an error. Note that you can pass an empty string to clear the input field.
		/// </para>
		/// <para>To send fine-grained keyboard events, use <see cref="Page.Type"/>.</para>
		/// <para>Shortcut for main frame's <see cref="Frame.Fill"/></para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="value">
		/// Value to fill for the <c>&lt;input&gt;</c>, <c>&lt;textarea&gt;</c> or <c>[contenteditable]</c>
		/// element.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Fill(*string selector, *string value, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method fetches an element with <paramref name="selector"/> and focuses it.
		/// If there's no element matching <paramref name="selector"/>, the method waits until
		/// a matching element appears in the DOM.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Focus"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Focus(*string selector, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Returns frame matching the specified criteria. Either <c>name</c> or <c>url</c>
		/// must be specified.
		/// </para>
		/// </summary>
		Frame Frame { get; }
		/// <summary><para>An array of all frames attached to the page.</para></summary>
		IReadOnlyCollection<Frame> Frames { get; }
		/// <summary><para>Returns element attribute value.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="name">Attribute name to get the value for.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		GetAttribute(*string selector, *string name, *float64 timeout = default) *string
		/// <summary>
		/// <para>
		/// Returns the main resource response. In case of multiple redirects, the navigation
		/// will resolve with the response of the last redirect. If can not go back, returns
		/// <c>null</c>.
		/// </para>
		/// <para>Navigate to the previous page in history.</para>
		/// </summary>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		GoBack(*float64 timeout = default, *WaitUntilState waitUntil = default) Response
		/// <summary>
		/// <para>
		/// Returns the main resource response. In case of multiple redirects, the navigation
		/// will resolve with the response of the last redirect. If can not go forward, returns
		/// <c>null</c>.
		/// </para>
		/// <para>Navigate to the next page in history.</para>
		/// </summary>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		GoForward(*float64 timeout = default, *WaitUntilState waitUntil = default) Response
		/// <summary>
		/// <para>
		/// Returns the main resource response. In case of multiple redirects, the navigation
		/// will resolve with the response of the last redirect.
		/// </para>
		/// <para><c>page.goto</c> will throw an error if:</para>
		/// <list type="bullet">
		/// <item><description>there's an SSL error (e.g. in case of self-signed certificates).</description></item>
		/// <item><description>target URL is invalid.</description></item>
		/// <item><description>the <paramref name="timeout"/> is exceeded during navigation.</description></item>
		/// <item><description>the remote server does not respond or is unreachable.</description></item>
		/// <item><description>the main resource failed to load.</description></item>
		/// </list>
		/// <para>
		/// <c>page.goto</c> will not throw an error when any valid HTTP status code is returned
		/// by the remote server, including 404 "Not Found" and 500 "Internal Server Error".
		/// The status code for such responses can be retrieved by calling <see cref="Response.Status"/>.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Goto"/></para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <c>page.goto</c> either throws an error or returns a main resource response. The
		/// only exceptions are navigation to <c>about:blank</c> or navigation to the same URL
		/// with a different hash, which would succeed and return <c>null</c>.
		/// </para>
		/// <para>
		/// Headless mode doesn't support navigation to a PDF document. See the <a href="https://bugs.chromium.org/p/chromium/issues/detail?id=761295">upstream
		/// issue</a>.
		/// </para>
		/// </remarks>
		/// <param name="url">URL to navigate page to. The url should include scheme, e.g. <c>https://</c>.</param>
		/// <param name="referer">
		/// Referer header value. If provided it will take preference over the referer header
		/// value set by <see cref="Page.SetExtraHttpHeaders"/>.
		/// </param>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		Goto(*string url, *string referer = default, *float64 timeout = default, *WaitUntilState waitUntil = default) Response
		/// <summary>
		/// <para>
		/// This method hovers over an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Mouse"/> to hover over the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <c>noWaitAfter</c>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Hover"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Hover(*string selector, *bool force = default, []KeyboardModifier modifiers = default, *PagePosition position = default, *float64 timeout = default) void
		/// <summary><para>Returns <c>element.innerHTML</c>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		InnerHTML(*string selector, *float64 timeout = default) *string
		/// <summary><para>Returns <c>element.innerText</c>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		InnerText(*string selector, *float64 timeout = default) *string
		/// <summary>
		/// <para>
		/// Returns whether the element is checked. Throws if the element is not a checkbox
		/// or radio input.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsChecked(*string selector, *float64 timeout = default) *bool
		/// <summary><para>Indicates that the page has been closed.</para></summary>
		*bool IsClosed { get; }
		/// <summary><para>Returns whether the element is disabled, the opposite of <a href="./actionability.md#enabled">enabled</a>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsDisabled(*string selector, *float64 timeout = default) *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#editable">editable</a>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsEditable(*string selector, *float64 timeout = default) *bool
		/// <summary><para>Returns whether the element is <a href="./actionability.md#enabled">enabled</a>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsEnabled(*string selector, *float64 timeout = default) *bool
		/// <summary>
		/// <para>
		/// Returns whether the element is hidden, the opposite of <a href="./actionability.md#visible">visible</a>.
		/// <paramref name="selector"/> that does not match any elements is considered hidden.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsHidden(*string selector, *float64 timeout = default) *bool
		/// <summary>
		/// <para>
		/// Returns whether the element is <a href="./actionability.md#visible">visible</a>.
		/// <paramref name="selector"/> that does not match any elements is considered not visible.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		IsVisible(*string selector, *float64 timeout = default) *bool
		Keyboard Keyboard `json:"keyboard"`
		/// <summary>
		/// <para>
		/// The page's main frame. Page is guaranteed to have a main frame which persists during
		/// navigations.
		/// </para>
		/// </summary>
		Frame MainFrame { get; }
		Mouse Mouse `json:"mouse"`
		/// <summary>
		/// <para>
		/// Returns the opener for popup pages and <c>null</c> for others. If the opener has
		/// been closed already the returns <c>null</c>.
		/// </para>
		/// </summary>
		GetOpener() Page
		/// <summary>
		/// <para>
		/// Pauses script execution. Playwright will stop executing the script and wait for
		/// the user to either press 'Resume' button in the page overlay or to call <c>playwright.resume()</c>
		/// in the DevTools console.
		/// </para>
		/// <para>
		/// User can inspect selectors or perform manual steps while paused. Resume will continue
		/// running the original script from the place it was paused.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// This method requires Playwright to be started in a headed mode, with a falsy <paramref
		/// name="headless"/> value in the <see cref="BrowserType.Launch"/>.
		/// </para>
		/// </remarks>
		Pause() void
		/// <summary>
		/// <para>Returns the PDF buffer.</para>
		/// <para>
		/// <c>page.pdf()</c> generates a pdf of the page with <c>print</c> css media. To generate
		/// a pdf with <c>screen</c> media, call <see cref="Page.EmulateMedia"/> before calling
		/// <c>page.pdf()</c>:
		/// </para>
		/// <para>
		/// The <paramref name="width"/>, <paramref name="height"/>, and <paramref name="margin"/>
		/// options accept values labeled with units. Unlabeled values are treated as pixels.
		/// </para>
		/// <para>A few examples:</para>
		/// <list type="bullet">
		/// <item><description><c>page.pdf({width: 100})</c> - prints with width set to 100 pixels</description></item>
		/// <item><description><c>page.pdf({width: '100px'})</c> - prints with width set to 100 pixels</description></item>
		/// <item><description><c>page.pdf({width: '10cm'})</c> - prints with width set to 10 centimeters.</description></item>
		/// </list>
		/// <para>All possible units are:</para>
		/// <list type="bullet">
		/// <item><description><c>px</c> - pixel</description></item>
		/// <item><description><c>in</c> - inch</description></item>
		/// <item><description><c>cm</c> - centimeter</description></item>
		/// <item><description><c>mm</c> - millimeter</description></item>
		/// </list>
		/// <para>The <paramref name="format"/> options are:</para>
		/// <list type="bullet">
		/// <item><description><c>Letter</c>: 8.5in x 11in</description></item>
		/// <item><description><c>Legal</c>: 8.5in x 14in</description></item>
		/// <item><description><c>Tabloid</c>: 11in x 17in</description></item>
		/// <item><description><c>Ledger</c>: 17in x 11in</description></item>
		/// <item><description><c>A0</c>: 33.1in x 46.8in</description></item>
		/// <item><description><c>A1</c>: 23.4in x 33.1in</description></item>
		/// <item><description><c>A2</c>: 16.54in x 23.4in</description></item>
		/// <item><description><c>A3</c>: 11.7in x 16.54in</description></item>
		/// <item><description><c>A4</c>: 8.27in x 11.7in</description></item>
		/// <item><description><c>A5</c>: 5.83in x 8.27in</description></item>
		/// <item><description><c>A6</c>: 4.13in x 5.83in</description></item>
		/// </list>
		/// </summary>
		/// <remarks>
		/// <para>Generating a pdf is currently only supported in Chromium headless.</para>
		/// <para>
		/// By default, <c>page.pdf()</c> generates a pdf with modified colors for printing.
		/// Use the <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/-webkit-print-color-adjust"><c>-webkit-print-color-adjust</c></a>
		/// property to force rendering of exact colors.
		/// </para>
		/// <para>
		/// <paramref name="headerTemplate"/> and <paramref name="footerTemplate"/> markup have
		/// the following limitations: > 1. Script tags inside templates are not evaluated.
		/// > 2. Page styles are not visible inside templates.
		/// </para>
		/// </remarks>
		/// <param name="displayHeaderFooter">Display header and footer. Defaults to <c>false</c>.</param>
		/// <param name="footerTemplate">
		/// HTML template for the print footer. Should use the same format as the <paramref
		/// name="headerTemplate"/>.
		/// </param>
		/// <param name="format">
		/// Paper format. If set, takes priority over <paramref name="width"/> or <paramref
		/// name="height"/> options. Defaults to 'Letter'.
		/// </param>
		/// <param name="headerTemplate">
		/// HTML template for the print header. Should be valid HTML markup with following classes
		/// used to inject printing values into them:
		/// <list type="bullet">
		/// <item><description><c>'date'</c> formatted print date</description></item>
		/// <item><description><c>'title'</c> document title</description></item>
		/// <item><description><c>'url'</c> document location</description></item>
		/// <item><description><c>'pageNumber'</c> current page number</description></item>
		/// <item><description><c>'totalPages'</c> total pages in the document</description></item>
		/// </list>
		/// </param>
		/// <param name="landscape">Paper orientation. Defaults to <c>false</c>.</param>
		/// <param name="pageRanges">
		/// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which
		/// means print all pages.
		/// </param>
		/// <param name="path">
		/// The file path to save the PDF to. If <paramref name="path"/> is a relative path,
		/// then it is resolved relative to the current working directory. If no path is provided,
		/// the PDF won't be saved to the disk.
		/// </param>
		/// <param name="preferCSSPageSize">
		/// Give any CSS <c>@page</c> size declared in the page priority over what is declared
		/// in <paramref name="width"/> and <paramref name="height"/> or <paramref name="format"/>
		/// options. Defaults to <c>false</c>, which will scale the content to fit the paper
		/// size.
		/// </param>
		/// <param name="printBackground">Print background graphics. Defaults to <c>false</c>.</param>
		/// <param name="scale">
		/// Scale of the webpage rendering. Defaults to <c>1</c>. Scale amount must be between
		/// 0.1 and 2.
		/// </param>
		Pdf(*bool displayHeaderFooter = default, *string footerTemplate = default, *string format = default, *string headerTemplate = default, *bool landscape = default, *string pageRanges = default, *string path = default, *bool preferCSSPageSize = default, *bool printBackground = default, *float64 scale = default) []byte
		/// <summary>
		/// <para>Focuses the element, and then uses <see cref="Keyboard.Down"/> and <see cref="Keyboard.Up"/>.</para>
		/// <para>
		/// <paramref name="key"/> can specify the intended <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key">keyboardEvent.key</a>
		/// value or a single character to generate the text for. A superset of the <paramref
		/// name="key"/> values can be found <a href="https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values">here</a>.
		/// Examples of the keys are:
		/// </para>
		/// <para>
		/// <c>F1</c> - <c>F12</c>, <c>Digit0</c>- <c>Digit9</c>, <c>KeyA</c>- <c>KeyZ</c>,
		/// <c>Backquote</c>, <c>Minus</c>, <c>Equal</c>, <c>Backslash</c>, <c>Backspace</c>,
		/// <c>Tab</c>, <c>Delete</c>, <c>Escape</c>, <c>ArrowDown</c>, <c>End</c>, <c>Enter</c>,
		/// <c>Home</c>, <c>Insert</c>, <c>PageDown</c>, <c>PageUp</c>, <c>ArrowRight</c>, <c>ArrowUp</c>,
		/// etc.
		/// </para>
		/// <para>
		/// Following modification shortcuts are also supported: <c>Shift</c>, <c>Control</c>,
		/// <c>Alt</c>, <c>Meta</c>, <c>ShiftLeft</c>.
		/// </para>
		/// <para>
		/// Holding down <c>Shift</c> will type the text that corresponds to the <paramref name="key"/>
		/// in the upper case.
		/// </para>
		/// <para>
		/// If <paramref name="key"/> is a single character, it is case-sensitive, so the values
		/// <c>a</c> and <c>A</c> will generate different respective texts.
		/// </para>
		/// <para>
		/// Shortcuts such as <c>key: "Control+o"</c> or <c>key: "Control+Shift+T"</c> are supported
		/// as well. When speficied with the modifier, modifier is pressed and being held while
		/// the subsequent key is being pressed.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="key">
		/// Name of the key to press or a character to generate, such as <c>ArrowLeft</c> or
		/// <c>a</c>.
		/// </param>
		/// <param name="delay">
		/// Time to wait between <c>keydown</c> and <c>keyup</c> in milliseconds. Defaults to
		/// 0.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Press(*string selector, *string key, *float64 delay = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// The method finds an element matching the specified selector within the page. If
		/// no elements match the selector, the return value resolves to <c>null</c>.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.QuerySelector"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		QuerySelector(*string selector) ElementHandle
		/// <summary>
		/// <para>
		/// The method finds all elements matching the specified selector within the page. If
		/// no elements match the selector, the return value resolves to <c>[]</c>.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.QuerySelectorAll"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		QuerySelectorAll(*string selector) IReadOnlyCollection<ElementHandle>
		/// <summary>
		/// <para>
		/// Returns the main resource response. In case of multiple redirects, the navigation
		/// will resolve with the response of the last redirect.
		/// </para>
		/// </summary>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		Reload(*float64 timeout = default, *WaitUntilState waitUntil = default) Response
		/// <summary>
		/// <para>Routing provides the capability to modify network requests that are made by a page.</para>
		/// <para>
		/// Once routing is enabled, every request matching the url pattern will stall unless
		/// it's continued, fulfilled or aborted.
		/// </para>
		/// <para>An example of a naïve handler that aborts all image requests:</para>
		/// <para>or the same snippet using a regex pattern instead:</para>
		/// <para>
		/// Page routes take precedence over browser context routes (set up with <see cref="BrowserContext.Route"/>)
		/// when request matches both handlers.
		/// </para>
		/// </summary>
		/// <remarks>
		/// <para>The handler will only be called for the first url if the response is a redirect.</para>
		/// <para>Enabling routing disables http cache.</para>
		/// </remarks>
		/// <param name="url">
		/// A glob pattern, regex pattern or predicate receiving <see cref="URL"/> to match
		/// while routing.
		/// </param>
		/// <param name="handler">handler function to route the request.</param>
		Route(interface{} url, func(Route, Request) handler) void
		/// <summary><para>Returns the buffer with the captured screenshot.</para></summary>
		/// <remarks>
		/// <para>
		/// Screenshots take at least 1/6 second on Chromium OS X and Chromium Windows. See
		/// https://crbug.com/741689 for discussion.
		/// </para>
		/// </remarks>
		/// <param name="clip">
		/// An object which specifies clipping of the resulting image. Should have the following
		/// fields:
		/// </param>
		/// <param name="fullPage">
		/// When true, takes a screenshot of the full scrollable page, instead of the currently
		/// visible viewport. Defaults to <c>false</c>.
		/// </param>
		/// <param name="omitBackground">
		/// Hides default white background and allows capturing screenshots with transparency.
		/// Not applicable to <c>jpeg</c> images. Defaults to <c>false</c>.
		/// </param>
		/// <param name="path">
		/// The file path to save the image to. The screenshot type will be inferred from file
		/// extension. If <paramref name="path"/> is a relative path, then it is resolved relative
		/// to the current working directory. If no path is provided, the image won't be saved
		/// to the disk.
		/// </param>
		/// <param name="quality">The quality of the image, between 0-100. Not applicable to <c>png</c> images.</param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="type">Specify screenshot type, defaults to <c>png</c>.</param>
		Screenshot(*PageClip clip = default, *bool fullPage = default, *bool omitBackground = default, *string path = default, *int quality = default, *float64 timeout = default, *ScreenshotType type = default) []byte
		/// <summary>
		/// <para>Returns the array of option values that have been successfully selected.</para>
		/// <para>
		/// Triggers a <c>change</c> and <c>input</c> event once all the provided options have
		/// been selected. If there's no <c>&lt;select&gt;</c> element matching <paramref name="selector"/>,
		/// the method throws an error.
		/// </para>
		/// <para>Will wait until all specified options are present in the <c>&lt;select&gt;</c> element.</para>
		/// <para>Shortcut for main frame's <see cref="Frame.SelectOption"/></para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SelectOption(*string selector, *bool noWaitAfter = default, *float64 timeout = default) IReadOnlyCollection<*string>
		/// <param name="html">HTML markup to assign to the page.</param>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		SetContent(*string html, *float64 timeout = default, *WaitUntilState waitUntil = default) void
		/// <summary>
		/// <para>
		/// This setting will change the default maximum navigation time for the following methods
		/// and related shortcuts:
		/// </para>
		/// <list type="bullet">
		/// <item><description><see cref="Page.GoBack"/></description></item>
		/// <item><description><see cref="Page.GoForward"/></description></item>
		/// <item><description><see cref="Page.Goto"/></description></item>
		/// <item><description><see cref="Page.Reload"/></description></item>
		/// <item><description><see cref="Page.SetContent"/></description></item>
		/// <item><description><see cref="Page.WaitForNavigation"/></description></item>
		/// </list>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <see cref="Page.SetDefaultNavigationTimeout"/> takes priority over <see cref="Page.SetDefaultTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/> and <see cref="BrowserContext.SetDefaultNavigationTimeout"/>.
		/// </para>
		/// </remarks>
		*float64 DefaultNavigationTimeout { set; }
		/// <summary>
		/// <para>
		/// This setting will change the default maximum time for all the methods accepting
		/// <paramref name="timeout"/> option.
		/// </para>
		/// </summary>
		/// <remarks><para><see cref="Page.SetDefaultNavigationTimeout"/> takes priority over <see cref="Page.SetDefaultTimeout"/>.</para></remarks>
		*float64 DefaultTimeout { set; }
		/// <summary><para>The extra HTTP headers will be sent with every request the page initiates.</para></summary>
		/// <remarks>
		/// <para>
		/// <see cref="Page.SetExtraHttpHeaders"/> does not guarantee the order of headers in
		/// the outgoing requests.
		/// </para>
		/// </remarks>
		/// <param name="headers">
		/// An object containing additional HTTP headers to be sent with every request. All
		/// header values must be strings.
		/// </param>
		SetExtraHttpHeaders(map[string]string headers) void
		/// <summary>
		/// <para>
		/// This method expects <paramref name="selector"/> to point to an <a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input">input
		/// element</a>.
		/// </para>
		/// <para>
		/// Sets the value of the file input to these file paths or files. If some of the <c>filePaths</c>
		/// are relative paths, then they are resolved relative to the the current working directory.
		/// For empty array, clears the selected files.
		/// </para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="files">
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		SetInputFiles(*string selector, []string files, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// In the case of multiple pages in a single browser, each page can have its own viewport
		/// size. However, <see cref="Browser.NewContext"/> allows to set viewport size (and
		/// more) for all pages in the context at once.
		/// </para>
		/// <para>
		/// <c>page.setViewportSize</c> will resize the page. A lot of websites don't expect
		/// phones to change size, so you should set the viewport size before navigating to
		/// the page.
		/// </para>
		/// </summary>
		SetViewportSize() void
		/// <summary>
		/// <para>
		/// This method taps an element matching <paramref name="selector"/> by performing the
		/// following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>
		/// Use <see cref="Page.Touchscreen"/> to tap the center of the element, or the specified
		/// <paramref name="position"/>.
		/// </description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Tap"/>.</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// <see cref="Page.Tap"/> requires that the <paramref name="hasTouch"/> option of the
		/// browser context be set to true.
		/// </para>
		/// </remarks>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="modifiers">
		/// Modifier keys to press. Ensures that only these modifiers are pressed during the
		/// operation, and then restores current modifiers back. If not specified, currently
		/// pressed modifiers are used.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="position">
		/// A point to use relative to the top-left corner of element padding box. If not specified,
		/// uses some visible point of the element.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Tap(*string selector, *bool force = default, []KeyboardModifier modifiers = default, *bool noWaitAfter = default, *PagePosition position = default, *float64 timeout = default) void
		/// <summary><para>Returns <c>element.textContent</c>.</para></summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		TextContent(*string selector, *float64 timeout = default) *string
		/// <summary><para>Returns the page's title. Shortcut for main frame's <see cref="Frame.Title"/>.</para></summary>
		GetTitle() *string
		Touchscreen Touchscreen `json:"touchscreen"`
		/// <summary>
		/// <para>
		/// Sends a <c>keydown</c>, <c>keypress</c>/<c>input</c>, and <c>keyup</c> event for
		/// each character in the text. <c>page.type</c> can be used to send fine-grained keyboard
		/// events. To fill values in form fields, use <see cref="Page.Fill"/>.
		/// </para>
		/// <para>To press a special key, like <c>Control</c> or <c>ArrowDown</c>, use <see cref="Keyboard.Press"/>.</para>
		/// <para>Shortcut for main frame's <see cref="Frame.Type"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="text">A text to type into a focused element.</param>
		/// <param name="delay">Time to wait between key presses in milliseconds. Defaults to 0.</param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Type(*string selector, *string text, *float64 delay = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// This method unchecks an element matching <paramref name="selector"/> by performing
		/// the following steps:
		/// </para>
		/// <list type="ordinal">
		/// <item><description>
		/// Find an element match matching <paramref name="selector"/>. If there is none, wait
		/// until a matching element is attached to the DOM.
		/// </description></item>
		/// <item><description>
		/// Ensure that matched element is a checkbox or a radio input. If not, this method
		/// rejects. If the element is already unchecked, this method returns immediately.
		/// </description></item>
		/// <item><description>
		/// Wait for <a href="./actionability.md">actionability</a> checks on the matched element,
		/// unless <paramref name="force"/> option is set. If the element is detached during
		/// the checks, the whole action is retried.
		/// </description></item>
		/// <item><description>Scroll the element into view if needed.</description></item>
		/// <item><description>Use <see cref="Page.Mouse"/> to click in the center of the element.</description></item>
		/// <item><description>
		/// Wait for initiated navigations to either succeed or fail, unless <paramref name="noWaitAfter"/>
		/// option is set.
		/// </description></item>
		/// <item><description>Ensure that the element is now unchecked. If not, this method rejects.</description></item>
		/// </list>
		/// <para>
		/// When all steps combined have not finished during the specified <paramref name="timeout"/>,
		/// this method rejects with a <see cref="TimeoutError"/>. Passing zero timeout disables
		/// this.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.Uncheck"/>.</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to search for element. If there are multiple elements satisfying the
		/// selector, the first will be used. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="force">
		/// Whether to bypass the <a href="./actionability.md">actionability</a> checks. Defaults
		/// to <c>false</c>.
		/// </param>
		/// <param name="noWaitAfter">
		/// Actions that initiate navigations are waiting for these navigations to happen and
		/// for pages to start loading. You can opt out of waiting via setting this flag. You
		/// would only need this option in the exceptional cases such as navigating to inaccessible
		/// pages. Defaults to <c>false</c>.
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		Uncheck(*string selector, *bool force = default, *bool noWaitAfter = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Removes a route created with <see cref="Page.Route"/>. When <paramref name="handler"/>
		/// is not specified, removes all routes for the <paramref name="url"/>.
		/// </para>
		/// </summary>
		/// <param name="url">
		/// A glob pattern, regex pattern or predicate receiving <see cref="URL"/> to match
		/// while routing.
		/// </param>
		/// <param name="handler">Optional handler function to route the request.</param>
		Unroute(interface{} url, func(Route, Request) handler = default) void
		/// <summary><para>Shortcut for main frame's <see cref="Frame.URL"/>.</para></summary>
		*string URL { get; }
		/// <summary><para>Video object associated with this page.</para></summary>
		Video Video { get; }
		*PageViewportSizeResult ViewportSize { get; }
		/// <summary>
		/// <para>
		/// Returns when the <paramref name="expression"/> returns a truthy value. It resolves
		/// to a JSHandle of the truthy value.
		/// </para>
		/// <para>The <see cref="Page.WaitForFunction"/> can be used to observe viewport size change:</para>
		/// <para>To pass an argument to the predicate of <see cref="Page.WaitForFunction"/> function:</para>
		/// <para>Shortcut for main frame's <see cref="Frame.WaitForFunction"/>.</para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		/// <param name="polling">
		/// If <paramref name="polling"/> is <c>'raf'</c>, then <paramref name="expression"/>
		/// is constantly executed in <c>requestAnimationFrame</c> callback. If <paramref name="polling"/>
		/// is a number, then it is treated as an interval in milliseconds at which the function
		/// would be executed. Defaults to <c>raf</c>.
		/// </param>
		/// <param name="timeout">
		/// maximum time to wait for in milliseconds. Defaults to <c>30000</c> (30 seconds).
		/// Pass <c>0</c> to disable timeout. The default value can be changed by using the
		/// <see cref="BrowserContext.SetDefaultTimeout"/>.
		/// </param>
		WaitForFunction(*string expression, interface{} arg = default, interface{} polling = default, *float64 timeout = default) JSHandle
		/// <summary>
		/// <para>Returns when the required load state has been reached.</para>
		/// <para>
		/// This resolves when the page reaches a required load state, <c>load</c> by default.
		/// The navigation must have been committed when this method is called. If current document
		/// has already reached the required state, resolves immediately.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.WaitForLoadState"/>.</para>
		/// </summary>
		/// <param name="state">
		/// Optional load state to wait for, defaults to <c>load</c>. If the state has been
		/// already reached while loading current document, the method resolves immediately.
		/// Can be one of:
		/// <list type="bullet">
		/// <item><description><c>'load'</c> - wait for the <c>load</c> event to be fired.</description></item>
		/// <item><description><c>'domcontentloaded'</c> - wait for the <c>DOMContentLoaded</c> event to be fired.</description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - wait until there are no network connections for at least
		/// <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForLoadState(*LoadState state = default, *float64 timeout = default) void
		/// <summary>
		/// <para>
		/// Waits for the main frame navigation and returns the main resource response. In case
		/// of multiple redirects, the navigation will resolve with the response of the last
		/// redirect. In case of navigation to a different anchor or navigation due to History
		/// API usage, the navigation will resolve with <c>null</c>.
		/// </para>
		/// <para>
		/// This resolves when the page navigates to a new URL or reloads. It is useful for
		/// when you run code which will indirectly cause the page to navigate. e.g. The click
		/// target has an <c>onclick</c> handler that triggers navigation from a <c>setTimeout</c>.
		/// Consider this example:
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.WaitForNavigation"/>.</para>
		/// </summary>
		/// <remarks>
		/// <para>
		/// Usage of the <a href="https://developer.mozilla.org/en-US/docs/Web/API/History_API">History
		/// API</a> to change the URL is considered a navigation.
		/// </para>
		/// </remarks>
		/// <param name="timeout">
		/// Maximum operation time in milliseconds, defaults to 30 seconds, pass <c>0</c> to
		/// disable timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultNavigationTimeout"/>,
		/// <see cref="BrowserContext.SetDefaultTimeout"/>, <see cref="Page.SetDefaultNavigationTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		/// <param name="url">
		/// A glob pattern, regex pattern or predicate receiving <see cref="URL"/> to match
		/// while waiting for the navigation.
		/// </param>
		/// <param name="waitUntil">
		/// When to consider operation succeeded, defaults to <c>load</c>. Events can be either:
		/// <list type="bullet">
		/// <item><description>
		/// <c>'domcontentloaded'</c> - consider operation to be finished when the <c>DOMContentLoaded</c>
		/// event is fired.
		/// </description></item>
		/// <item><description>
		/// <c>'load'</c> - consider operation to be finished when the <c>load</c> event is
		/// fired.
		/// </description></item>
		/// <item><description>
		/// <c>'networkidle'</c> - consider operation to be finished when there are no network
		/// connections for at least <c>500</c> ms.
		/// </description></item>
		/// </list>
		/// </param>
		WaitForNavigation(*float64 timeout = default, interface{} url = default, *WaitUntilState waitUntil = default) Response
		/// <summary><para>Waits for the matching request and returns it.</para></summary>
		/// <param name="urlOrPredicate*string">Request URL string, regex or predicate receiving <see cref="Request"/> object.</param>
		/// <param name="urlOrPredicateRegex">Request URL string, regex or predicate receiving <see cref="Request"/> object.</param>
		/// <param name="urlOrPredicateFunc">Request URL string, regex or predicate receiving <see cref="Request"/> object.</param>
		/// <param name="timeout">
		/// Maximum wait time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable
		/// the timeout. The default value can be changed by using the <see cref="Page.SetDefaultTimeout"/>
		/// method.
		/// </param>
		WaitForRequest(*string urlOrPredicate*string, Regex urlOrPredicateRegex, Func<Request, *bool> urlOrPredicateFunc, *float64 timeout = default) Request
		/// <summary><para>Returns the matched response.</para></summary>
		/// <param name="urlOrPredicate*string">Request URL string, regex or predicate receiving <see cref="Response"/> object.</param>
		/// <param name="urlOrPredicateRegex">Request URL string, regex or predicate receiving <see cref="Response"/> object.</param>
		/// <param name="urlOrPredicateFunc">Request URL string, regex or predicate receiving <see cref="Response"/> object.</param>
		/// <param name="timeout">
		/// Maximum wait time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable
		/// the timeout. The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForResponse(*string urlOrPredicate*string, Regex urlOrPredicateRegex, Func<Response, *bool> urlOrPredicateFunc, *float64 timeout = default) Response
		/// <summary>
		/// <para>
		/// Returns when element specified by selector satisfies <paramref name="state"/> option.
		/// Returns <c>null</c> if waiting for <c>hidden</c> or <c>detached</c>.
		/// </para>
		/// <para>
		/// Wait for the <paramref name="selector"/> to satisfy <paramref name="state"/> option
		/// (either appear/disappear from dom, or become visible/hidden). If at the moment of
		/// calling the method <paramref name="selector"/> already satisfies the condition,
		/// the method will return immediately. If the selector doesn't satisfy the condition
		/// for the <paramref name="timeout"/> milliseconds, the function will throw.
		/// </para>
		/// <para>This method works across navigations:</para>
		/// </summary>
		/// <param name="selector">
		/// A selector to query for. See <a href="./selectors.md">working with selectors</a>
		/// for more details.
		/// </param>
		/// <param name="state">
		/// Defaults to <c>'visible'</c>. Can be either:
		/// <list type="bullet">
		/// <item><description><c>'attached'</c> - wait for element to be present in DOM.</description></item>
		/// <item><description><c>'detached'</c> - wait for element to not be present in DOM.</description></item>
		/// <item><description>
		/// <c>'visible'</c> - wait for element to have non-empty bounding box and no <c>visibility:hidden</c>.
		/// Note that element without any content or with <c>display:none</c> has an empty bounding
		/// box and is not considered visible.
		/// </description></item>
		/// <item><description>
		/// <c>'hidden'</c> - wait for element to be either detached from DOM, or have an empty
		/// bounding box or <c>visibility:hidden</c>. This is opposite to the <c>'visible'</c>
		/// option.
		/// </description></item>
		/// </list>
		/// </param>
		/// <param name="timeout">
		/// Maximum time in milliseconds, defaults to 30 seconds, pass <c>0</c> to disable timeout.
		/// The default value can be changed by using the <see cref="BrowserContext.SetDefaultTimeout"/>
		/// or <see cref="Page.SetDefaultTimeout"/> methods.
		/// </param>
		WaitForSelector(*string selector, *WaitForSelectorState state = default, *float64 timeout = default) ElementHandle
		/// <summary>
		/// <para>Waits for the given <paramref name="timeout"/> in milliseconds.</para>
		/// <para>
		/// Note that <c>page.waitForTimeout()</c> should only be used for debugging. Tests
		/// using the timer in production are going to be flaky. Use signals such as network
		/// events, selectors becoming visible and others instead.
		/// </para>
		/// <para>Shortcut for main frame's <see cref="Frame.WaitForTimeout"/>.</para>
		/// </summary>
		/// <param name="timeout">A timeout to wait for</param>
		WaitForTimeout(*float64 timeout) void
		/// <summary>
		/// <para>
		/// This method returns all of the dedicated <a href="https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API">WebWorkers</a>
		/// associated with the page.
		/// </para>
		/// </summary>
		/// <remarks><para>This does not contain ServiceWorkers</para></remarks>
		IReadOnlyCollection<Worker> Workers { get; }
	}/// <summary>
	/// <para>
	/// Playwright module provides a method to launch a browser instance. The following
	/// is a typical example of using Playwright to drive automation:
	/// </para>
	/// </summary>
	type Playwright interface {
		/// <summary>
		/// <para>
		/// This object can be used to launch or connect to Chromium, returning instances of
		/// <see cref="ChromiumBrowser"/>.
		/// </para>
		/// </summary>
		Chromium BrowserType `json:"chromium"`
		/// <summary>
		/// <para>
		/// This object can be used to launch or connect to Firefox, returning instances of
		/// <see cref="FirefoxBrowser"/>.
		/// </para>
		/// </summary>
		Firefox BrowserType `json:"firefox"`
		/// <summary>
		/// <para>
		/// Selectors can be used to install custom selector engines. See <a href="./selectors.md">Working
		/// with selectors</a> for more information.
		/// </para>
		/// </summary>
		Selectors Selectors `json:"selectors"`
		/// <summary>
		/// <para>
		/// This object can be used to launch or connect to WebKit, returning instances of <see
		/// cref="WebKitBrowser"/>.
		/// </para>
		/// </summary>
		Webkit BrowserType `json:"webkit"`
	}/// <summary>
	/// <para>
	/// Whenever the page sends a request for a network resource the following sequence
	/// of events are emitted by <see cref="Page"/>:
	/// </para>
	/// <list type="bullet">
	/// <item><description><see cref="Page.Request"/> emitted when the request is issued by the page.</description></item>
	/// <item><description>
	/// <see cref="Page.Response"/> emitted when/if the response status and headers are
	/// received for the request.
	/// </description></item>
	/// <item><description>
	/// <see cref="Page.RequestFinished"/> emitted when the response body is downloaded
	/// and the request is complete.
	/// </description></item>
	/// </list>
	/// <para>
	/// If request fails at some point, then instead of <c>'requestfinished'</c> event (and
	/// possibly instead of 'response' event), the  <see cref="Page.RequestFailed"/> event
	/// is emitted.
	/// </para>
	/// <para>
	/// If request gets a 'redirect' response, the request is successfully finished with
	/// the 'requestfinished' event, and a new request is  issued to a redirected url.
	/// </para>
	/// </summary>
	/// <remarks>
	/// <para>
	/// HTTP Error responses, such as 404 or 503, are still successful responses from HTTP
	/// standpoint, so request will complete with <c>'requestfinished'</c> event.
	/// </para>
	/// </remarks>
	type Request interface {
		/// <summary>
		/// <para>
		/// The method returns <c>null</c> unless this request has failed, as reported by <c>requestfailed</c>
		/// event.
		/// </para>
		/// <para>Example of logging of all the failed requests:</para>
		/// </summary>
		*string Failure { get; }
		/// <summary><para>Returns the <see cref="Frame"/> that initiated this request.</para></summary>
		Frame Frame { get; }
		/// <summary><para>An object with HTTP headers associated with the request. All header names are lower-case.</para></summary>
		map[string]string Headers { get; }
		/// <summary><para>Whether this request is driving frame's navigation.</para></summary>
		*bool IsNavigationRequest { get; }
		/// <summary><para>Request's method (GET, POST, etc.)</para></summary>
		*string Method { get; }
		/// <summary><para>Request's post body, if any.</para></summary>
		*string PostData { get; }
		/// <summary><para>Request's post body in a binary form, if any.</para></summary>
		[]byte PostDataBuffer { get; }
		/// <summary>
		/// <para>Request that was redirected by the server to this one, if any.</para>
		/// <para>
		/// When the server responds with a redirect, Playwright creates a new <see cref="Request"/>
		/// object. The two requests are connected by <c>redirectedFrom()</c> and <c>redirectedTo()</c>
		/// methods. When multiple server redirects has happened, it is possible to construct
		/// the whole redirect chain by repeatedly calling <c>redirectedFrom()</c>.
		/// </para>
		/// <para>For example, if the website <c>http://example.com</c> redirects to <c>https://example.com</c>:</para>
		/// <para>If the website <c>https://google.com</c> has no redirects:</para>
		/// </summary>
		Request RedirectedFrom { get; }
		/// <summary>
		/// <para>New request issued by the browser if the server responded with redirect.</para>
		/// <para>This method is the opposite of <see cref="Request.RedirectedFrom"/>:</para>
		/// </summary>
		Request RedirectedTo { get; }
		/// <summary>
		/// <para>
		/// Contains the request's resource type as it was perceived by the rendering engine.
		/// ResourceType will be one of the following: <c>document</c>, <c>stylesheet</c>, <c>image</c>,
		/// <c>media</c>, <c>font</c>, <c>script</c>, <c>texttrack</c>, <c>xhr</c>, <c>fetch</c>,
		/// <c>eventsource</c>, <c>websocket</c>, <c>manifest</c>, <c>other</c>.
		/// </para>
		/// </summary>
		*string ResourceType { get; }
		/// <summary>
		/// <para>
		/// Returns the matching <see cref="Response"/> object, or <c>null</c> if the response
		/// was not received due to error.
		/// </para>
		/// </summary>
		GetResponse() Response
		/// <summary>
		/// <para>
		/// Returns resource timing information for given request. Most of the timing values
		/// become available upon the response, <c>responseEnd</c> becomes available when request
		/// finishes. Find more information at <a href="https://developer.mozilla.org/en-US/docs/Web/API/PerformanceResourceTiming">Resource
		/// Timing API</a>.
		/// </para>
		/// </summary>
		*RequestTimingResult Timing { get; }
		/// <summary><para>URL of the request.</para></summary>
		*string URL { get; }
	}/// <summary><para><see cref="Response"/> class represents responses which are received by page.</para></summary>
	type Response interface {
		/// <summary><para>Returns the buffer with response body.</para></summary>
		GetBody() []byte
		/// <summary><para>Waits for this response to finish, returns failure error if request failed.</para></summary>
		GetFinished() *string
		/// <summary><para>Returns the <see cref="Frame"/> that initiated this response.</para></summary>
		Frame Frame { get; }
		/// <summary>
		/// <para>
		/// Returns the object with HTTP headers associated with the response. All header names
		/// are lower-case.
		/// </para>
		/// </summary>
		map[string]string Headers { get; }
		/// <summary>
		/// <para>
		/// Contains a boolean stating whether the response was successful (status in the range
		/// 200-299) or not.
		/// </para>
		/// </summary>
		*bool Ok { get; }
		/// <summary><para>Returns the matching <see cref="Request"/> object.</para></summary>
		Request Request { get; }
		/// <summary><para>Contains the status code of the response (e.g., 200 for a success).</para></summary>
		*int Status { get; }
		/// <summary><para>Contains the status text of the response (e.g. usually an "OK" for a success).</para></summary>
		*string StatusText { get; }
		/// <summary><para>Returns the text representation of response body.</para></summary>
		GetText() *string
		/// <summary><para>Contains the URL of the response.</para></summary>
		*string URL { get; }
	}/// <summary>
	/// <para>
	/// Whenever a network route is set up with <see cref="Page.Route"/> or <see cref="BrowserContext.Route"/>,
	/// the <c>Route</c> object allows to handle the route.
	/// </para>
	/// </summary>
	type Route interface {
		/// <summary><para>Aborts the route's request.</para></summary>
		/// <param name="errorCode">
		/// Optional error code. Defaults to <c>failed</c>, could be one of the following:
		/// <list type="bullet">
		/// <item><description><c>'aborted'</c> - An operation was aborted (due to user action)</description></item>
		/// <item><description>
		/// <c>'accessdenied'</c> - Permission to access a resource, other than the network,
		/// was denied
		/// </description></item>
		/// <item><description>
		/// <c>'addressunreachable'</c> - The IP address is unreachable. This usually means
		/// that there is no route to the specified host or network.
		/// </description></item>
		/// <item><description><c>'blockedbyclient'</c> - The client chose to block the request.</description></item>
		/// <item><description>
		/// <c>'blockedbyresponse'</c> - The request failed because the response was delivered
		/// along with requirements which are not met ('X-Frame-Options' and 'Content-Security-Policy'
		/// ancestor checks, for instance).
		/// </description></item>
		/// <item><description>
		/// <c>'connectionaborted'</c> - A connection timed out as a result of not receiving
		/// an ACK for data sent.
		/// </description></item>
		/// <item><description><c>'connectionclosed'</c> - A connection was closed (corresponding to a TCP FIN).</description></item>
		/// <item><description><c>'connectionfailed'</c> - A connection attempt failed.</description></item>
		/// <item><description><c>'connectionrefused'</c> - A connection attempt was refused.</description></item>
		/// <item><description><c>'connectionreset'</c> - A connection was reset (corresponding to a TCP RST).</description></item>
		/// <item><description><c>'internetdisconnected'</c> - The Internet connection has been lost.</description></item>
		/// <item><description><c>'namenotresolved'</c> - The host name could not be resolved.</description></item>
		/// <item><description><c>'timedout'</c> - An operation timed out.</description></item>
		/// <item><description><c>'failed'</c> - A generic failure occurred.</description></item>
		/// </list>
		/// </param>
		Abort(*string errorCode = default) void
		/// <summary><para>Continues route's request with optional overrides.</para></summary>
		/// <param name="headers">If set changes the request HTTP headers. Header values will be converted to a string.</param>
		/// <param name="method">If set changes the request method (e.g. GET or POST)</param>
		/// <param name="postData">If set changes the post data of request</param>
		/// <param name="url">If set changes the request URL. New URL must have same protocol as original one.</param>
		Continue(map[string]string headers = default, *string method = default, interface{} postData = default, *string url = default) void
		/// <summary>
		/// <para>Fulfills route's request with given response.</para>
		/// <para>An example of fulfilling all requests with 404 responses:</para>
		/// <para>An example of serving static file:</para>
		/// </summary>
		/// <param name="body">Response body.</param>
		/// <param name="contentType">If set, equals to setting <c>Content-Type</c> response header.</param>
		/// <param name="headers">Response headers. Header values will be converted to a string.</param>
		/// <param name="path">
		/// File path to respond with. The content type will be inferred from file extension.
		/// If <c>path</c> is a relative path, then it is resolved relative to the current working
		/// directory.
		/// </param>
		/// <param name="status">Response status code, defaults to <c>200</c>.</param>
		Fulfill(interface{} body = default, *string contentType = default, map[string]string headers = default, *string path = default, *int status = default) void
		/// <summary><para>A request to be routed.</para></summary>
		Request Request { get; }
	}/// <summary>
	/// <para>
	/// Selectors can be used to install custom selector engines. See <a href="./selectors.md">Working
	/// with selectors</a> for more information.
	/// </para>
	/// </summary>
	type Selectors interface {
		/// <summary><para>An example of registering selector engine that queries elements based on a tag name:</para></summary>
		/// <param name="name">
		/// Name that is used in selectors as a prefix, e.g. <c>{name: 'foo'}</c> enables <c>foo=myselectorbody</c>
		/// selectors. May only contain <c>[a-zA-Z0-9_]</c> characters.
		/// </param>
		/// <param name="contentScript">
		/// Whether to run this selector engine in isolated JavaScript environment. This environment
		/// has access to the same DOM, but not any JavaScript objects from the frame's scripts.
		/// Defaults to <c>false</c>. Note that running as a content script is not guaranteed
		/// when this engine is used together with other registered engines.
		/// </param>
		Register(*string name, *bool contentScript = default) void
	}/// <summary>
	/// <para>
	/// The Touchscreen class operates in main-frame CSS pixels relative to the top-left
	/// corner of the viewport. Methods on the touchscreen can only be used in browser contexts
	/// that have been intialized with <c>hasTouch</c> set to true.
	/// </para>
	/// </summary>
	type Touchscreen interface {
		/// <summary>
		/// <para>
		/// Dispatches a <c>touchstart</c> and <c>touchend</c> event with a single touch at
		/// the position (<paramref name="x"/>,<paramref name="y"/>).
		/// </para>
		/// </summary>
		/// <param name="x">
		/// </param>
		/// <param name="y">
		/// </param>
		Tap(*float64 x, *float64 y) void
	}/// <summary>
	/// <para>
	/// When browser context is created with the <c>videosPath</c> option, each page has
	/// a video object associated with it.
	/// </para>
	/// </summary>
	type Video interface {
		/// <summary>
		/// <para>
		/// Returns the file system path this video will be recorded to. The video is guaranteed
		/// to be written to the filesystem upon closing the browser context.
		/// </para>
		/// </summary>
		GetPath() *string
	}/// <summary><para>The <see cref="WebSocket"/> class represents websocket connections in the page.</para></summary>
	type WebSocket interface {
		/// <summary><para>Fired when the websocket closes.</para></summary>
		event EventHandler Close;
		/// <summary><para>Fired when the websocket recieves a frame.</para></summary>
		event EventHandler<*FrameReceivedPayload> FrameReceived;
		/// <summary><para>Fired when the websocket sends a frame.</para></summary>
		event EventHandler<*FrameSentPayload> FrameSent;
		/// <summary><para>Fired when the websocket has an error.</para></summary>
		event EventHandler<String> SocketError;
		/// <summary><para>Indicates that the web socket has been closed.</para></summary>
		*bool IsClosed { get; }
		/// <summary><para>Contains the URL of the WebSocket.</para></summary>
		*string URL { get; }
	}/// <summary>
	/// <para>
	/// The Worker class represents a <a href="https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API">WebWorker</a>.
	/// <c>worker</c> event is emitted on the page object to signal a worker creation. <c>close</c>
	/// event is emitted on the worker object when the worker is gone.
	/// </para>
	/// <code>
	/// Page.Worker += (_, worker) =&gt;<br/>
	/// {<br/>
	///     Console.WriteLine($"Worker created: {worker.Url}");<br/>
	///     worker.Close += (_, _) =&gt; Console.WriteLine($"Worker closed {worker.Url}");<br/>
	/// };<br/>
	/// <br/>
	/// Console.WriteLine("Current Workers:");<br/>
	/// foreach(var pageWorker in Page.Workers)<br/>
	/// {<br/>
	///     Console.WriteLine($"\tWorker: {pageWorker.Url}");<br/>
	/// }
	/// </code>
	/// </summary>
	type Worker interface {
		/// <summary>
		/// <para>
		/// Emitted when this dedicated <a href="https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API">WebWorker</a>
		/// is terminated.
		/// </para>
		/// </summary>
		event EventHandler Close;
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/>.</para>
		/// <para>
		/// If the function passed to the <see cref="Worker.Evaluate"/> returns a <see cref="Promise"/>,
		/// then <see cref="Worker.Evaluate"/> would wait for the promise to resolve and return
		/// its value.
		/// </para>
		/// <para>
		/// If the function passed to the <see cref="Worker.Evaluate"/> returns a non-<see cref="Serializable"/>
		/// value, then <see cref="Worker.Evaluate"/> returns <c>undefined</c>. Playwright also
		/// supports transferring some additional values that are not serializable by <c>JSON</c>:
		/// <c>-0</c>, <c>NaN</c>, <c>Infinity</c>, <c>-Infinity</c>.
		/// </para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		Evaluate(*string expression, interface{} arg = default) interface{}
		/// <summary>
		/// <para>Returns the return value of <paramref name="expression"/> as a <see cref="JSHandle"/>.</para>
		/// <para>
		/// The only difference between <see cref="Worker.Evaluate"/> and <see cref="Worker.EvaluateHandle"/>
		/// is that <see cref="Worker.EvaluateHandle"/> returns <see cref="JSHandle"/>.
		/// </para>
		/// <para>
		/// If the function passed to the <see cref="Worker.EvaluateHandle"/> returns a <see
		/// cref="Promise"/>, then <see cref="Worker.EvaluateHandle"/> would wait for the promise
		/// to resolve and return its value.
		/// </para>
		/// </summary>
		/// <param name="expression">
		/// JavaScript expression to be evaluated in the browser context. If it looks like a
		/// function declaration, it is interpreted as a function. Otherwise, evaluated as an
		/// expression.
		/// </param>
		/// <param name="arg">Optional argument to pass to <paramref name="expression"/>.</param>
		EvaluateHandle(*string expression, interface{} arg = default) JSHandle
		*string URL { get; }
	}