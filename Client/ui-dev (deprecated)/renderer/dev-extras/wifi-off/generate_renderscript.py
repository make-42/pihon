from PIL import Image


img = Image.open("./pihon-banner.png")

screenSize = [128,64]
screenCount = 2

imgload = img.load()

finalOutput = ""

currentlyDrawingLine = False
startPosition = [0,0]

def generate_script_string(startPosition,length):
    return "DrawHorizontalLine("+str(startPosition[0])+", "+str(startPosition[1])+", "+str(length)+", 1)\n"
	

for y in range(img.size[1]):
    for x in range(img.size[0]):
        if imgload[x,y] == (255,255,255,255):
            if currentlyDrawingLine == False:
                currentlyDrawingLine = True
                startPosition = [x,y]
        else:
            if currentlyDrawingLine == True:
                currentlyDrawingLine = False
                if x == 0:
                    finalOutput += generate_script_string(startPosition,img.size[0]-1-startPosition[0])
                else:
                    finalOutput += generate_script_string(startPosition,x-startPosition[0])
            
print(finalOutput)