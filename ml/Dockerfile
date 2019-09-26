FROM python:3.7.4

ENV WORKDIR /app/
WORKDIR ${WORKDIR}

COPY Pipfile Pipfile.lock ${WORKDIR}

RUN pip install pipenv --no-cache-dir && \
    pipenv install --system --deploy && \
    pip uninstall -y pipenv virtualenv-clone virtualenv

COPY . ${WORKDIR}

USER nobody
CMD ["python", "main.py"]