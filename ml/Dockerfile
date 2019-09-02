FROM python:3.7.4


WORKDIR /app
COPY . .

RUN pip install -r requirements.txt --upgrade
USER nobody

CMD ["python3", "main.py"]