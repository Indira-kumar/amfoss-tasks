import argparse
from urllib import response
import requests

parser = argparse.ArgumentParser(description="Enter date and id of image to be fetched as in-line arguments")

parser.add_argument("date", type=str, help="Enter date in YYYY-MM-DD form")
parser.add_argument("id",type=int, help="Enter id of the image to be filtered out")

args = parser.parse_args()
date = args.date
imgid = args.id

URL ="https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?earth_date="+date+"&api_key=ba0dxv6ebQHn3zcVEtKrVeseF2zVzW6wzzhkazab"
response = requests.get(URL)
jsonFile = response.json()

for i in jsonFile["photos"]:
    if(i['id']==imgid):
        print(i['img_src'])
        break
else:
    print("No Match Found")