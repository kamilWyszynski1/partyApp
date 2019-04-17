package service.email.registration;

import org.apache.tomcat.util.http.parser.HttpParser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.mail.MessagingException;
import javax.persistence.EntityNotFoundException;
import java.util.List;
import java.util.Optional;
import java.util.Random;


@RestController
public class RegisterCodeController {

    private static final String REDIS_INDEX_KEY = "PRODUCT";
//
//    @Autowired
//    RedisTemplate<String, Object> redisTemplate;

    @Autowired
    private RegisterCodeRepository registerCodeRepository;

    @Autowired
    private EmailServiceImpl emailService;

    @RequestMapping(value = "/verification", method = RequestMethod.POST)
    public RegisterCode createRegisterCode(@RequestBody RegisterCode registerCode) {
        int code = new Random().nextInt(900000) + 100000;
        registerCode.setCode(code);

        try {
            emailService.sendMessage("kamil-wyszynski@wp.pl", String.valueOf(code));
        } catch (MessagingException e) {
            e.printStackTrace();
        }

        registerCodeRepository.save(registerCode);
        return registerCode;
    }

    @RequestMapping(value = "/verification", method = RequestMethod.GET)
    public List<RegisterCode> getCodes(){
        return (List<RegisterCode>) registerCodeRepository.findAll();
    }

    @RequestMapping(value = "/verification", method = RequestMethod.DELETE)
    public void deleteCodes(){
        registerCodeRepository.deleteAll();
    }

    @RequestMapping(value = "/verify", method = RequestMethod.POST)
    public ResponseEntity<RegisterCode> verify(@RequestBody RegisterCode registerCode){

        RegisterCode registerCodeDb = registerCodeRepository.findById(registerCode.getEmail())
                .orElse(null);

        if (registerCodeDb == null)
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);

        if (registerCodeDb.getCode() == registerCode.getCode())
            return new ResponseEntity<>(HttpStatus.OK);
        else
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
    }
}
