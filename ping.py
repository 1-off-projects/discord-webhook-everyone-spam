from discord import SyncWebhook
import time, discord, random, string, threading
# your webhook
hookurl = ('your webhook')
webhook = SyncWebhook.from_url(hookurl)
while True:    
    webhook.send('@everyone')