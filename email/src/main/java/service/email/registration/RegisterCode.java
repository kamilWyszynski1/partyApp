package service.email.registration;

import org.springframework.data.annotation.Id;
import org.springframework.data.redis.core.RedisHash;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;

@RedisHash("registerCode")
public class RegisterCode {
    public RegisterCode() {
    }

    public RegisterCode(String email) {
        this.email = email;
    }

    private int code;

    @Id
    private String email;


    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}

