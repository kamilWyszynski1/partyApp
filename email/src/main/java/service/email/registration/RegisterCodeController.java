package service.email.registration;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import javax.mail.MessagingException;
import java.util.List;
import java.util.Random;

@CrossOrigin(origins = "http://localhost:8080")
@RestController
public class RegisterCodeController {

    @Value("${access_token}")
    private String access_token;

    private final RegisterCodeRepository registerCodeRepository;

    private final EmailServiceImpl emailService;

    public RegisterCodeController(RegisterCodeRepository registerCodeRepository, EmailServiceImpl emailService) {
        this.registerCodeRepository = registerCodeRepository;
        this.emailService = emailService;
    }

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
        /*
            Endpoint verify if sent code matches email, if so it sends request to activate
            user to User Service and return STATUS 200. If Code doesnt match email it returns
            404 or 401.
         */

        RegisterCode registerCodeDb = registerCodeRepository.findById(registerCode.getEmail())
                .orElse(null);

        if (registerCodeDb == null)
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);

        if (registerCodeDb.getCode() == registerCode.getCode()) {
            // Delete verified email
            registerCodeRepository.delete(registerCode);
            this.activate_user(registerCode.getEmail());

            return new ResponseEntity<>(HttpStatus.OK);
        }
        else
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
    }

    private void activate_user(String email){
        /*
            Function sends request to User Service and call acitvate user function
         */
        final String url = "http://localhost:8000/activate/";
        RestTemplate restTemplate = new RestTemplate();
        String result = restTemplate.getForObject(url, String.class);

        System.out.println(result);

    }
}
