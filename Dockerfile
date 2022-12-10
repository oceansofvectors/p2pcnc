# Start from a base image with the necessary tools
FROM golang:1.15

# Copy the binary file to the root directory
COPY filename /

# Make the binary executable
RUN chmod +x /filename

# Set the binary as the entrypoint for the container
ENTRYPOINT ["/filename"]
