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
    uses: th0th/notify-discord@v0.4.1
    if: ${{ always() }}
    env:
      DISCORD_WEBHOOK_URL: ${{ secrets.DISCORD_WEBHOOK_URL }}
      GITHUB_ACTOR: ${{ github.actor }}
      GITHUB_JOB_NAME: "Build and deploy"
      GITHUB_JOB_STATUS: ${{ job.status }}
```

## Shameless plug

I am an indie hacker, and I am running two services that might be useful for your business. Check them out :)

### WebGazer

[<img alt="WebGazer" src="https://user-images.githubusercontent.com/698079/162474223-f7e819c4-4421-4715-b8a2-819583550036.png" width="256" />](https://www.webgazer.io/?utm_source=github&utm_campaign=postgres-s3-backup-readme)

WebGazer is a monitoring service that checks your website, cron jobs, or scheduled tasks on a regular basis. It notifies
you with instant alerts in case of a problem. That way, you have peace of mind about the status of your service without
manually checking it.

### PoeticMetric

[<img alt="PoeticMetric" src="https://user-images.githubusercontent.com/698079/162474946-7c4565ba-5097-4a42-8821-d087e6f56a5d.png" width="256" />](https://www.poeticmetric.com/?utm_source=github&utm_campaign=postgres-s3-backup-readme)

PoeticMetric is a privacy-first, regulation-compliant, blazingly fast analytics tool.

No cookies or personal data collection. So you don't have to worry about cookie banners or GDPR, CCPA, and PECR compliance.

## License

Copyright © 2021, Gökhan Sarı. Released under the [MIT License](LICENSE).
