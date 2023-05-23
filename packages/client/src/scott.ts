import { fetch } from "node-fetch";

fetch("https://www.election2033.xyz/api/addMatchResults", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-US,en;q=0.9",
    "content-type": "application/json",
    "sec-ch-ua": "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin"
  },
  "referrer": "https://www.election2033.xyz/",
  "referrerPolicy": "strict-origin-when-cross-origin",
  "body": "{\"user_id\":\"94a50edf-d476-441f-8073-5f55bac03b24\",\"opponent_id\":\"ce81e5b2-4946-4ca5-8c69-55a899e22985\",\"winner\":\"ese_efekodo\"}",
  "method": "POST",
  "mode": "cors",
  "credentials": "include"
});