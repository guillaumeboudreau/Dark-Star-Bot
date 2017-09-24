# Dark-Star-Bot
Role management and server management discord bot made for the Dark Star Community, all in golang.

Using the discordgo library [https://www.github.com/bwmarrin/discordgo]

# To compile:

* Make sure that you have your gopath set and you have added bwmarrin/discordgo from Github to it, otherwise it will not compile.

1) Download the Go compiler from http://golang.org

2) Go to the location of the source files

3) Type `go build` [I'd advise against `go run` because you have to use a token to run the bot correctly.]

4) Continue to running the bot.

# To run the bot for yourself:

1) Make a discord app

2) Make the discord app a bot

3) Take the token from the app and put it in a .bat or a .sh like this: Dark-Star-Bot.exe -t TOKENHERE

4) Use an OAuth2 Generator to get an invite link

5) MAKE SURE that you have a 'roles/' directory with the application, and that you have a .dsr file! That is how the bot knows which roles you can call from the bot. See below to see how to make a .dsr file properly.

6) Invite to your server and wammo!

# How to make a .dsr file:

## You need a serverconfig.dsr file! Please add that first.

It's structured like so: [No spaces!]

GuildID=roles/FileName.dsr;

GuildID2=roles/FileName2.dsr;

* Example:

123123123123=roles/myServer.dsr;

123131231231=roles/myFriendsServer.dsr;

## .DSR server file:

* It's structured like so:

WHEN MAKING THE .DSR FILES, YOU MUST FOLLOW THIS FORMAT: [PLEASE READ]

LINE 1: Always the server ID with a ; on the end

LINE 2: Always the welcome Channel ID with a ; on the end

LINE 3: Always the welcome Channel Message with a ; on the end,

     Add a %s where the username should go! [Ex: Welcome, %s!]

LINE 4: Always the goodbye Channel ID with a ; on the end

LINE 5: Always the goodbye Channel Message with a ; on the end

     Add a %s where the username should go! [Ex: Welcome, %s!]

LINE 6: Always the ban     Channel ID with a ; on the end

LINE 7: Always the ban     Channel Message with a ; on the end

     Add a %s where the username should go! [Ex: Welcome, %s!]

LINES PROCEEDING: Calls=Role;

     Example: call1,call2=Role;

* The calls MUST NOT have spaces, the Roles CAN have spaces

exampleone=Example One;

## A call is what the user will type to request a role,

* Example:
#### IN DSR FILE:

GuildID;

welcomeChannelID;

Hello, welcome to the server, %s!;

goodbyeChannelID;

Goodbye, %s!;

banChannelID;

Get out, %s!;

role1,roleone=Role1;

role2,roletwo=Role2;

#### USER CAN TYPE THESE TO CALL THE ROLE:

'$role role1'

'$role roleone'

#### Make SURE that the .dsr files are in a directory called 'roles' and that directory is in the same directory as the application.

* Example:

~/Documents/Folder1/Application

~/Documents/Folder1/roles/serverconfig.dsr

~/Documents/Folder1/roles/server1roles.dsr
