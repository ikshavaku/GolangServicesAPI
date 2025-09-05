-- +goose Up
-- +goose StatementBegin
INSERT INTO public.services
(id, "name", description, created_at, updated_at, is_deleted)
VALUES(1, 'test_1', 'dummy', '2025-08-31 18:27:53.228', '2025-08-31 18:27:53.228', false);
INSERT INTO public.services
(id, "name", description, created_at, updated_at, is_deleted)
VALUES(2, 'service1', 'dummy', '2025-08-31 18:27:53.228', '2025-08-31 18:27:53.228', false);
INSERT INTO public.services
(id, "name", description, created_at, updated_at, is_deleted)
VALUES(3, 'service2', 'dummy', '2025-08-31 18:27:53.228', '2025-08-31 18:27:53.228', false);
INSERT INTO public.services
(id, "name", description, created_at, updated_at, is_deleted)
VALUES(4, 'service4', 'dummy', '2025-08-31 18:27:53.228', '2025-08-31 18:27:53.228', false);
INSERT INTO public.services
(id, "name", description, created_at, updated_at, is_deleted)
VALUES(5, 'service5', 'dummy', '2025-08-31 18:27:53.228', '2025-08-31 18:27:53.228', false);
INSERT INTO public.service_versions
(service_id, version_number, created_at, updated_at, is_deleted)
VALUES(1, '1', '2025-08-31 18:28:46.346', '2025-08-31 18:28:46.346', false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM public.service_versions
WHERE service_id=1 AND version_number='1';
DELETE FROM public.services
WHERE id=1;
DELETE FROM public.services
WHERE id=2;
DELETE FROM public.services
WHERE id=3;
DELETE FROM public.services
WHERE id=4;
DELETE FROM public.services
WHERE id=5;
-- +goose StatementEnd
