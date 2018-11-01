1. 定時に発火してLambda1が実行されること
2. 完了後、Lambda2がレスポンスfalseを返すまで1分おきに繰り返し実行されること

```
{
  "StartAt": "Lambda1",
  "States": {
    "Lambda1": {
      "Type": "Task",
      "Resource": "arn:aws:lambda:ap-northeast-1:{AWS_ACCOUNT_ID}:function:step-functions-example_lambda1",
      "Next": "Wait1"
    },
    "Wait1": {
      "Type": "Wait",
      "Seconds": 10,
      "Next": "Lambda2"
    },
    "Lambda2": {
      "Type": "Task",
      "Resource": "arn:aws:lambda:ap-northeast-1:{AWS_ACCOUNT_ID}:function:step-functions-example_lambda2",
      "ResultPath": "$.iterator",
      "Next": "CheckLoop"
    },
    "CheckLoop": {
      "Type": "Choice",
      "Choices": [{
        "Variable": "$.iterator.next_loop",
        "BooleanEquals": true,
        "Next": "Wait1"
      }],
      "Default": "ExitProcess"
    },
    "ExitProcess": {
      "Type": "Succeed"
    }
  }
}
```

