from datetime import datetime, date
startDate = date(1989, 2, 12)
while startDate.year < 2020:
    startDate = startDate.replace(startDate.year + 1)
    print(startDate.strftime("%A, %d. %B %Y"))
