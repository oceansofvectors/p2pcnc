# Start from a base image with the necessary tools
FROM golang:1.15

# Copy the binary file to the root directory
COPY server /

# Make the binary executable
RUN chmod +x /server
