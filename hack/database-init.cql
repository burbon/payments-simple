CREATE KEYSPACE IF NOT EXISTS paymentssimple
WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor' : 1};

CREATE TABLE IF NOT EXISTS paymentssimple.payments (
    id UUID,
    type text,
    organisation_id UUID,
    version text,
    attributes text,
    PRIMARY KEY(id)
);
