import ebooklib
import re
from ebooklib import epub

filenames = ["Convenience Store Woman (Sayaka Murata) (z-lib.org)","The Martian_ A Novel - Weir, Andy"]
input_folder="./input/"
output_folder="./output/"

# Code
for filename in filenames:
    book = epub.read_epub(input_folder+filename+".epub")

    book_title = book.get_metadata('DC', 'title')[0][0]
    book_creator = book.get_metadata('DC', 'creator')[0][0]

    print(book_title+", "+book_creator)

    converted_book_content = ""

    for item in book.get_items():
        if item.get_type() == ebooklib.ITEM_DOCUMENT:
            converted_book_content+=item.get_body_content().decode("UTF-8")

    converted_book_content = converted_book_content.replace("<p class=\"crt1\">","")
    converted_book_content = converted_book_content.replace("<p class=\"crt\">","")
    converted_book_content = converted_book_content.replace("<p class=\"toc\">","")
    converted_book_content = converted_book_content.replace("<p class=\"indent\">", "  ")
    converted_book_content = converted_book_content.replace("<p class=\"nonindent\">","")
    converted_book_content = converted_book_content.replace("<p class=\"center\">", "    ")
    converted_book_content = converted_book_content.replace("<p class=\"extract\">", "")
    converted_book_content = converted_book_content.replace("</p>", "")

    converted_book_content = converted_book_content.replace("<div>", "")
    converted_book_content = converted_book_content.replace("</div>", "")

    converted_book_content = converted_book_content.replace("&amp;","&")
    converted_book_content = converted_book_content.replace("&#13;", "\r")
    converted_book_content = converted_book_content.replace("<br/>", "\n")

    converted_book_content = converted_book_content.replace("<em>","") # Ignore italics
    converted_book_content = converted_book_content.replace("</em>","") # Ignore italics
    converted_book_content = converted_book_content.replace("<strong>","") # Ignore bold text
    converted_book_content = converted_book_content.replace("</strong>","") # Ignore bold text

    converted_book_content = converted_book_content.replace("<span class=\"small\">","") # Ignore small text
    converted_book_content = converted_book_content.replace("<span class=\"dropcaps\">","")
    converted_book_content = converted_book_content.replace("</span>","")

    div = False
    header_a = False
    header_b = False
    header_a_open = False
    header_b_open = False
    final_converted_book_content = ""
    for character_index in range(len(converted_book_content)):
        # Opening tags
        if len(converted_book_content)-character_index > 3:
            if converted_book_content[character_index]+converted_book_content[character_index+1]+converted_book_content[character_index+2] == "<h1":
                header_a = True
                header_a_open = True
                final_converted_book_content += "# "
        
        if len(converted_book_content)-character_index > 3:
            if converted_book_content[character_index]+converted_book_content[character_index+1]+converted_book_content[character_index+2] == "<h2":
                header_b = True
                header_b_open = True
                final_converted_book_content += "## "

        if len(converted_book_content)-character_index > 4:
            if converted_book_content[character_index]+converted_book_content[character_index+1]+converted_book_content[character_index+2]+converted_book_content[character_index+3] == "<div":
                div = True
        # Closing tags
        if len(converted_book_content)-character_index > 4:
            if converted_book_content[character_index]+converted_book_content[character_index+1]+converted_book_content[character_index+2]+converted_book_content[character_index+3] == "</h1":
                header_a_open = False
        if len(converted_book_content)-character_index > 4:
            if converted_book_content[character_index]+converted_book_content[character_index+1]+converted_book_content[character_index+2]+converted_book_content[character_index+3] == "</h2":
                header_b_open = False
        # Check for end of opening tag
        if converted_book_content[character_index-1] == ">":
            if header_a:
                header_a = False
            if header_b:
                header_b = False
            if div:
                div = False
        # Remove opening tags
        if header_a:
            pass
        elif header_b:
            pass
        elif div:
            pass
        else:
            final_converted_book_content+=converted_book_content[character_index]
        # Add # to headers
        if converted_book_content[character_index] == "\n":
            if header_a_open:
                final_converted_book_content+="# "
            if header_b_open:
                final_converted_book_content+="## "
    
    # Cleanup unused data
    final_converted_book_content = final_converted_book_content.replace("</h1>","")
    final_converted_book_content = final_converted_book_content.replace("</h2>","")
    final_converted_book_content = final_converted_book_content.replace("</body>","")
    final_converted_book_content = final_converted_book_content.replace("</a>","")
    final_converted_book_content = final_converted_book_content.replace("</div>","")

    final_converted_book_content = re.sub("(?s)<body.*?>", "", final_converted_book_content)
    final_converted_book_content = re.sub("(?s)<p.*?>", "\n", final_converted_book_content)
    final_converted_book_content = re.sub("(?s)<span.*?>", "", final_converted_book_content)
    final_converted_book_content = re.sub("(?s)<img.*?/>", "", final_converted_book_content)
    final_converted_book_content = re.sub("(?s)<a.*?>", "", final_converted_book_content)
    final_converted_book_content = re.sub("(?s)<div.*?/>", "", final_converted_book_content)
    final_converted_book_content = re.sub("(?s)<div.*?>", "", final_converted_book_content)

    # Strip out return carriage line breaks
    final_converted_book_content = final_converted_book_content.replace("\r","")
    
    # Generate final_save_data
    final_save_data = ""

    final_save_data += "-+---+-BOOK TITLE-+---+-\n"
    final_save_data += book_title+"\n"
    final_save_data += "-+---+-BOOK AUTHOR-+---+-\n"
    final_save_data += book_creator+"\n"
    final_save_data += "-+---+-BOOK CONTENT-+---+-\n"
    final_save_data += final_converted_book_content

    output_file = open(output_folder+filename+".phn","w+")
    output_file.write(final_save_data)
    output_file.close()