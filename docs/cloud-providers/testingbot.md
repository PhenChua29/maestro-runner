# TestingBot

Run Maestro YAML flows on [TestingBot](https://testingbot.com)'s real device cloud via the Appium driver.

## Android

Create `testingbot-android.json`:

```json
{
  "platformName": "Android",
  "appium:deviceName": "Pixel 8",
  "appium:platformVersion": "14.0",
  "appium:app": "tb://app-id",
  "appium:automationName": "UiAutomator2",
  "tb:options": {
    "key": "YOUR_TESTINGBOT_KEY",
    "secret": "YOUR_TESTINGBOT_SECRET",
    "name": "Maestro Android test"
  }
}
```

Run:

```bash
maestro-runner --driver appium \
  --appium-url "https://hub.testingbot.com/wd/hub" \
  --caps testingbot-android.json \
  test flows/
```

## iOS

Create `testingbot-ios.json`:

```json
{
  "platformName": "iOS",
  "appium:deviceName": "iPhone 16",
  "appium:platformVersion": "18.0",
  "appium:app": "tb://app-id",
  "appium:automationName": "XCUITest",
  "tb:options": {
    "key": "YOUR_TESTINGBOT_KEY",
    "secret": "YOUR_TESTINGBOT_SECRET",
    "name": "Maestro iOS test",
    "realDevice": true
  }
}
```

Run:

```bash
maestro-runner --driver appium \
  --appium-url "https://hub.testingbot.com/wd/hub" \
  --caps testingbot-ios.json \
  test flows/
```

## Uploading Your App

Upload your `.apk`, `.ipa`, or `.zip` to TestingBot before running tests:

```bash
curl -X POST "https://api.testingbot.com/v1/storage" \
  -u "YOUR_TESTINGBOT_KEY:YOUR_TESTINGBOT_SECRET" \
  -F "file=@/path/to/app.apk"
```

The response returns an `app_url` (e.g., `tb://app-id`) — use that as the `appium:app` value in your caps file.
