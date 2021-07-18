Webhook Error for Bugsnag to Telegram channel

Callback json from Bugsnag
```json
{
	"account": {
		"id": "554b23dc381929f412244e87",
		"name": "Test Account",
		"url": "http://app.bugsnag.com/accounts/test-account?i=wh&m=te"
	},
	"project": {
		"id": "559bc00a94a804b02a96dd81",
		"name": "Example.com",
		"url": "http://app.bugsnag.com/projects/example?i=wh&m=te"
	},
	"trigger": {
		"type": "exception",
		"message": "Test error"
	},
	"error": {
		"id": "54d535b2776562445f002e4f",
		"errorId": "530c3522e694aae2f4000035",
		"exceptionClass": "ExampleException",
		"message": "Something really bad happened",
		"context": "home#example",
		"appVersion": "App version",
		"releaseStage": "production",
		"firstReceived": "2014-02-25T06:16:02.000Z",
		"receivedAt": "2015-01-01T00:00:00.000Z",
		"userId": "554c1d7e5a724d0b36f327c1",
		"assignedUserId": "559c1d7e5a724d0b76f327c6",
		"url": "http://app.bugsnag.com/errors/example/events/example?i=wh&m=te",
		"severity": "error",
		"status": "open",
		"unhandled": true,
		"app": {
			"id": "com.bugsnag.android.example.debug",
			"version": "App version",
			"versionCode": "12",
			"bundleVersion": "1.0.2",
			"codeBundleId": "1.0-1234",
			"buildUUID": "BE5BA3D0-971C-4418-9ECF-E2D1ABCB66BE",
			"releaseStage": "production",
			"type": "rails",
			"dsymUUIDs": ["5E9D820E-3C31-364E-BE2D-8C319D17D8F6"],
			"duration": 1275,
			"durationInForeground": 983,
			"inForeground": true,
			"binaryArch": "x86_64",
			"runningOnRosetta": false
		},
		"device": {
			"hostname": "web1.internal",
			"id": "fd124e87760c4281aef",
			"manufacturer": "LGE",
			"model": "Nexus 6P",
			"modelNumber": "600",
			"osName": "android",
			"osVersion": "8.0.1",
			"freeMemory": 183879616,
			"totalMemory": 201326592,
			"freeDisk": 13478064128,
			"browserName": "Chrome",
			"browserVersion": "61.0.3163.100",
			"jailbroken": false,
			"orientation": "portrait",
			"locale": "en_US",
			"charging": false,
			"batteryLevel": 0.99,
			"time": "2015-01-15T02:42:16.000Z",
			"timezone": "PST",
			"macCatalystiOSVersion": "14.0"
		},
		"user": {
			"id": "554c1d7e5a724d0b36f327c1",
			"name": "John Doe",
			"email": "john@doe.com"
		},
		"stackTrace": [{
			"file": "app/controllers/home_controller.rb",
			"lineNumber": "123",
			"method": "example",
			"inProject": true,
			"code": {
				"121": "    crashy_function",
				"122": "  end",
				"123": "",
				"124": "  def crashy_function",
				"125": "    raise TestException.new(\"This is a test exception\")",
				"126": "  end",
				"127": "end"
			}
		}, {
			"file": "app/controllers/other_controller.rb",
			"lineNumber": "12",
			"method": "broken",
			"inProject": true
		}, {
			"file": "gems/junk/junkfile.rb",
			"lineNumber": "999",
			"method": "something",
			"inProject": false
		}, {
			"file": "lib/important/magic.rb",
			"lineNumber": "4",
			"method": "load_something",
			"inProject": true
		}],
		"breadcrumbs": []
	}
}

```