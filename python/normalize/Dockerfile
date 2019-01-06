# Use an official Python runtime as a parent image
FROM python:3.7.1-slim

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Create test directory structure
RUN bin/create-dirs.sh && \
    bin/create-files.sh

# Install any needed packages specified in requirements.txt
#RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Make port 80 available to the world outside this container
#EXPOSE 80

# Define environment variable
#ENV NAME World

# Entrypoint for application
Entrypoint ["python", "normalize/normalize.py"]

# Run app.py when the container launches
#CMD ["python", "normalize/normalize.py"]
