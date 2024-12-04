FROM node:20-alpine
WORKDIR /app
COPY . /app
RUN npm install serve -g
RUN npm install
RUN npm run build
EXPOSE 80
CMD ["npm", "run", "serve"]
