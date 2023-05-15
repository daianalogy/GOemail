# **Email sender to multiple recipients with Go**

My first golang program it purpose is to send emails to various recipients, you can add this Go script to Crontab in order to make email sending automatic.


## ***How to use ?***

This script requires the setting of some enviroments variables in order to function. These env vars are relate with our sender email. 

You **require to export** these enviroment variables correctly.

*This programm help you with that task the first time you run it*

You can set the values to the variables needed and then, the programm will give you the encoded value to export.

## ***How to export?*** 

Follow the instructions and Copy the variables line with its encoded value.

Then in your **Linux Enviroment** apply the following command.


> export [paste your enviroment variables and encoded values here]



## **Enviroment Variables**


| Variable | Description | Test Value |
| :---                |   :---:     |       ---: |
| **SET_SMTP** | set your email provider (gmail, outlook, yahoo) read the documentation to each one| smtp.example.com |
| **SET_EMAIL** | the email sender | user@example.com |
| **SET_PASS** | sender email's password | ********* |



## **About Security**

You can set email configuration trough  enviroment variables before run this programm for the first time but it's up to you if you set encoded values or not.

If you do it trought the first execution of this program your email configuration enviroment variables will be encoded. 


