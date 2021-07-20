`notify-discord` is a docker image that sends message to Discord using webhooks.

## Usage

I wrote this action to send a message to Discord when a GitHub action completes.

1. Create a webhook on your Discord server's settings.
1. Add the webhook URL to your repo's/organization's secrets with the name of `DISCORD_WEBHOOK_URL`. Or if you don't mind, you can directly paste the URL into the action file replacing `${{ secrets.DISCORD_WEBHOOK_URL }}`. 

By default, if a step fails, GitHub actions doesn't trigger the following ones. If you add `if: ${{ always() }}`, this action includes the status of the action in the message that is sent to Discord. 

```yaml
steps:
  - name: Your previous build steps
    
  - name: Notify discord
    uses: th0th/notify-discord@v0.3
    if: ${{ always() }}
    env:
      DISCORD_WEBHOOK_URL: ${{ secrets.DISCORD_WEBHOOK_URL }}
      GITHUB_ACTOR: ${{ github.actor }}
      GITHUB_JOB_STATUS: ${{ job.status }}
```

## Shameless plug

I am an indie hacker and I am running an uptime monitoring and analytics platform called [WebGazer](https://www.webgazer.io). You might want to check it out if you are running an online business and want to notice the incidents before your customers.

## License

Copyright © 2021, Gökhan Sarı. Released under the [MIT License](LICENSE).
