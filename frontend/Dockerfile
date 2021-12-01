FROM node:17-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

COPY . .
RUN npm i

# start app
CMD ["npm", "run", "start"]
