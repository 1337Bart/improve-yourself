import json
import sys
from openai import OpenAI
from dotenv import load_dotenv
import os

client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))


def load_json_file(file_path):
    try:
        with open(file_path, 'r') as file:
            return json.load(file)
    except FileNotFoundError:
        print(f"Error: File '{file_path}' not found")
        sys.exit(1)
    except json.JSONDecodeError:
        print(f"Error: File '{file_path}' is not valid JSON")
        sys.exit(1)

def call_openai_api(prompt, api_key):
    try:
        response = client.chat.completions.create(model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": prompt},
        ],
        n=1,
        stop=None,
        temperature=0.7)
        return response.choices[0].message.content.strip()
    except Exception as e:
        print(f"Error calling OpenAI API: {e}")
        sys.exit(1)


def main():
    load_dotenv()

    OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
    if len(sys.argv) != 2:
        print("Usage: python3 your_script.py <path_to_json_file>")
        sys.exit(1)

    json_file = sys.argv[1]
    data = load_json_file(json_file)

    prompt = \
        (f"Below you will find activity data per day over the last few days. "
         f"your job is to find ACTIONABLE INSIGHTS using the data provided. "
         f"Insights can be, for example: `on days with worse sleep score you were more likely to be less productive`. "
         f"or `you were happier on days with less screen time`, etc."
         f" your job is to find all possible correlations in data provided. "
         f"Use ONLY provided data to find your conclussions.  You should provide at least 6 insights."
         f"After providing insights, use them (DO NOT USE YOUR KNOWLEDGE FOR THIS - JUST THE INSIGHTS) to propose activities to improve one's life."
         f"activity data::\n{json.dumps(data, indent=4)}")

    reply = call_openai_api(prompt, OPENAI_API_KEY)

    print("##### python:  #######")
    print(reply)

if __name__ == "__main__":
    main()
