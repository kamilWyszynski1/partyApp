package service.email.registration;

import org.springframework.data.repository.CrudRepository;

public interface RegisterCodeRepository extends CrudRepository<RegisterCode, String> {
}
