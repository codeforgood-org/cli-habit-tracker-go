FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY habit /usr/local/bin/habit

# Create directory for habit data
RUN mkdir -p /root/.habit-tracker

# Set default environment variable
ENV HABIT_DATA_FILE=/root/.habit-tracker/habits.json

ENTRYPOINT ["habit"]
CMD ["help"]
