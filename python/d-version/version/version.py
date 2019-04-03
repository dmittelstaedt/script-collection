from bs4 import BeautifulSoup
import requests
import sys

version_url = "http://localhost:8080"
index_html = "my-index.txt"
search_string = "KIDICAP Version"


def get_old_version(input_file, search_string):
    """ Return the old version of the KIDICAP installation. The path to
    to the current index.html and a search string is used as arguments.
    """
    try:
        page_content = BeautifulSoup(open(input_file), "html.parser")
    except FileNotFoundError:
        return None
    version = page_content.find(class_="small")
    if version is not None:
        return version.text.strip()
    return None


def get_current_version(url_to_parse):
    """ Return the current version of the KIDICAP installation. The url to the
    KIDICAP.Comand index.html is used as argument.
    """
    try:
        page_response = requests.get(version_url)
    except requests.ConnectionError:
        return None
    page_content = BeautifulSoup(page_response.content, "html.parser")
    version = page_content.find(class_='versionInfo')
    if version is not None:
        return version.text.strip()
    return None


def write_new_version(output_file, old_version, new_version):
    """ Write the current version to the index.html of the overview website.
    This function needs the path to the output index.html, the old version of
    the KIDICAP installation and the new version of the KIDICAP version as
    arguments.
    """
    try:
        with open(output_file, "r") as file:
            content = file.read()
        with open(output_file, "w") as file:
            content = content.replace(old_version, new_version)
            file.write(content)
    except FileNotFoundError:
        print("Output file " + output_file + " could not be found.")


if __name__ == "__main__":
    old_version = get_old_version(index_html, search_string)
    new_version = get_current_version(version_url)
    if old_version is not None and new_version is not None:
        if old_version != new_version:
            write_new_version(index_html, old_version, new_version)
    else:
        print("Error during parsing input values.")
        sys.exit(1)
