# Use an official Node.js image
FROM node:20

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory
WORKDIR /app

# Copy package.json and pnpm-lock.yaml to take advantage of Docker cache
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install --frozen-lockfile --store-dir=.pnpm-store

# Copy the rest of the application files
COPY . .

# Expose the port for Vite
EXPOSE 5173

# Start the application
CMD ["pnpm", "dev"]
