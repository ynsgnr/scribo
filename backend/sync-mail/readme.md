# sync-email

Stateless, just sends whatever data in the request to given mail. Also polls email service and confirms any message from amazon kindle automatically

# Run:

docker run -e SMTP_EMAIL=$env:SMTP_EMAIL -e SMTP_PORT_EMAIL=$env:SMTP_PORT_EMAIL -e IMAP_EMAIL=$env:IMAP_EMAIL -e IMAP_PORT_EMAIL=$env:IMAP_PORT_EMAIL -e USERNAME_EMAIL=$env:USERNAME_EMAIL -e PASS_EMAIL=$env:PASS_EMAIL -e FROM_EMAIL=$env:FROM_EMAIL -e HOME="/home" -v $env:USERPROFILE\.aws:/home/.aws  --network=infrastructure_default -t sync-mail