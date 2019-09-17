from datetime import datetime, date
startDate = date(1989,2,12)
while startDate.year<2020:
 endDate = startDate.replace(startDate.year + 1)
 print(endDate.strftime("%A, %d. %B %Y"))
 startDate = startDate.replace(startDate.year + 1)
