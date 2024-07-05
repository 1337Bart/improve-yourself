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

    prompt = (
        "Below you will find activity data per day over the last few days. "
        "Your job is to find ACTIONABLE INSIGHTS using the data provided. "
        "Insights can be, for example: 'on days with worse sleep score you were more likely to be less productive'. "
        "or 'you were happier on days with less screen time', etc. "
        "Your job is to find all possible correlations in data provided. "
        "Use ONLY provided data to find your conclusions. You should provide at least 6 insights. "
        "After providing insights, use them (DO NOT USE YOUR KNOWLEDGE FOR THIS - JUST THE INSIGHTS) to propose activities to improve one's life. "
        "Below is the format for how the insights and proposed activities should be structured:\n"
        "{\n"
        "    \"insights\": [\n"
        "        {\"description\": \"Higher Sleep Score Correlates with Better Productivity\", \"details\": \"Days with higher sleep scores tend to have better productivity scores.\"},\n"
        "        {\"description\": \"Higher Social Interactions Lead to Higher Happiness Ratings\", \"details\": \"Days with more social interactions are associated with higher happiness ratings.\"},\n"
        "        {\"description\": \"Increased Outdoor Time Correlates with Lower Stress Levels\", \"details\": \"Days with more time spent outdoors tend to have lower stress levels.\"},\n"
        "        {\"description\": \"Higher Screen Time Hours Linked to Lower Sleep Scores\", \"details\": \"Days with more screen time hours often have lower sleep scores.\"},\n"
        "        {\"description\": \"More Meals Eaten Correlate with Higher Water Consumption\", \"details\": \"Days with more meals eaten typically show higher water consumption in liters.\"},\n"
        "        {\"description\": \"Higher Alcohol Consumption Associated with Lower Productivity Scores\", \"details\": \"Days with more alcohol units consumed might lead to lower productivity scores.\"}\n"
        "    ],\n"
        "    \"proposed_activities\": [\n"
        "        {\"description\": \"Improve Sleep Quality for Better Productivity\", \"details\": \"Establish a bedtime routine to improve sleep quality, leading to increased productivity.\"},\n"
        "        {\"description\": \"Enhance Social Interactions for Greater Happiness\", \"details\": \"Plan social activities or connect with friends and family regularly to boost happiness levels.\"},\n"
        "        {\"description\": \"Increase Outdoor Time for Stress Reduction\", \"details\": \"Incorporate outdoor activities like walks or outdoor workouts to reduce stress levels.\"},\n"
        "        {\"description\": \"Limit Screen Time for Improved Sleep\", \"details\": \"Set screen time limits before bedtime to improve sleep quality and overall health.\"},\n"
        "        {\"description\": \"Balanced Meal Planning for Hydration\", \"details\": \"Focus on consuming regular, balanced meals to ensure proper hydration throughout the day.\"},\n"
        "        {\"description\": \"Moderate Alcohol Consumption for Better Productivity\", \"details\": \"Monitor and moderate alcohol intake to maintain higher productivity levels and mental clarity.\"}\n"
        "    ]\n"
        "}\n"
        "Each section should have its content in bullet points and should be at least 6 points long. "
        f"Activity data::\n{json.dumps(data, indent=4)}"
    )

    reply = call_openai_api(prompt, OPENAI_API_KEY)

    print(reply)

if __name__ == "__main__":
    main()
