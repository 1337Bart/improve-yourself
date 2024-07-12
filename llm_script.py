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

    prompt = f"""
        Below you will find activity data per day over the last few days.
        Your task is to analyze the provided data and generate actionable insights based on the correlations you find.
        Examples of insights could be: 'On days with worse sleep scores, you were more likely to be less productive' or 'You were happier on days with less screen time'.
        Focus on finding all possible correlations within the given data. Use ONLY the provided data to draw your conclusions. You should provide at least 6 insights.
        After providing the insights, use them (WITHOUT USING ANY EXTERNAL KNOWLEDGE) to propose activities that could improve one's life based on the identified correlations.
        Please structure the insights and proposed activities in the following JSON format:
        
        {{
            "insights": [
                {{"description": "Insight 1 title", "details": "Insight 1 detailed explanation"}},
                {{"description": "Insight 2 title", "details": "Insight 2 detailed explanation"}},
                ...
            ],
            "proposed_activities": [
                {{"description": "Activity 1 title", "details": "Activity 1 detailed description"}},
                {{"description": "Activity 2 title", "details": "Activity 2 detailed description"}},
                ...
            ]
        }}
        
        Each section should contain at least 6 items, with the content presented in a clear and concise manner.
        Here is the activity data in JSON format:
            \\n{json.dumps(data, indent=4)}
        """



    reply = call_openai_api(prompt, OPENAI_API_KEY)

    print(reply)

if __name__ == "__main__":
    main()
