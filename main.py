from flask import Flask, render_template, request
from langchain_openai import OpenAI
import os

app = Flask(__name__)

OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")

def generate_response(input_text):
    try:
        llm = OpenAI(temperature=0.7, openai_api_key=OPENAI_API_KEY)
        return llm(input_text)
    except Exception as e:
        print(f"An error occurred: {e}")
        return None
try:
    print(generate_response("Are you there?"))
except Exception as e:
    print(f"An error occurred: {e}")
    
@app.route('/', methods=['GET', 'POST'])
def index():
    poem = None
    if request.method == 'POST':
        poem_title = request.form.get('poem_title', '')
        lang = request.form.get('language', '')
        if lang:
            prompt = f'''You are a very good and creative poem writer, write a poem for me on title "{poem_title}", generate this poem in "{lang}" language. Please note if you can't generate the poem on given topic or you can'nt understand the given topic then simply say "I can't do it" nothing else or any other excuse.'''
            poem = generate_response(prompt)
        else:
            prompt = f'''You are a very good and creative poem writer, write a poem for me on title "{poem_title}", generate this poem in "English" language.Please note if you can't generate the poem on given topic or you can'nt understand the given topic then simply say "I can't do it" nothing else or any other excuse.'''
            poem = generate_response(prompt)
    return render_template('index.html', poem=poem)
    
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80, debug=True)
