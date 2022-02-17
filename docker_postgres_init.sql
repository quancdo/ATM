CREATE TABLE accounts(
   id uuid DEFAULT uuid_generate_v4 (),
   name TEXT NOT NULL,
   api_key TEXT NOT NULL,
   balance INT NOT NULL
   pin TEXT NOT NULL
);

INSERT INTO accountss(name, api_key, balancer) VALUES
 ('John', 'john_apikey', 1000, "1234"),
 ('Michael', 'michael_apikey', 5, "5678"),
 ('Simon', 'simon_apikey', 3000, "1234");