# JingWei-Flowable-Plus - Agent Guide

## Build Commands

```bash
# Clean and compile
mvn clean compile

# Run all tests
mvn test

# Run specific test class
mvn test -Dtest=ClassName

# Run specific test method
mvn test -Dtest=ClassName#methodName

# Run tests with profile (local/dev/prod)
mvn test -Pdev

# Package application
mvn clean package

# Run application
java -jar jingwei-admin/target/jingwei-admin.jar

# Skip tests during build
mvn clean package -DskipTests
```

## Project Structure

Multi-module Maven project:
- `jingwei-admin` - Main application entry point and web controllers
- `jingwei-common` - Common utilities and base classes
- `jingwei-system` - System management (users, roles, depts, menus, dict)
- `jingwei-framework` - Framework configuration and core components
- `jingwei-flowable` - Workflow engine integration
- `jingwei-generator` - Code generator
- `jingwei-job` - Scheduled tasks (XXL-Job)
- `jingwei-oss` - Object storage service
- `jingwei-sms` - SMS service
- `jingwei-demo` - Demo module
- `jingwei-ztuo-guard` - Guard/security related business logic

## Code Style Guidelines

### Imports
- Group imports: standard java, third-party, project-specific
- Avoid wildcard imports except for static imports
- Use `javax.annotation.Resource` for dependency injection in tests
- Use `org.springframework.beans.factory.annotation.Autowired` for main code

### Formatting
- Use 4 spaces for indentation (no tabs)
- Line length: 120 characters preferred
- Class braces on same line
- Method braces on same line

### Types and Annotations
- Use `@RestController` for REST controllers
- Use `@Service` for service layer
- Use `@RequiredArgsConstructor` (Lombok) for constructor injection
- Use `@Validated` at class level for validation
- Use validation groups: `AddGroup.class`, `EditGroup.class`
- Use `@Slf4j` (Lombok) for logging

### Naming Conventions
- Controllers: `XxxController` (e.g., `SysUserController`)
- Services: `IXxxService` (interface), `XxxServiceImpl` (implementation)
- BOs (Business Objects): `XxxBo`
- VOs (View Objects): `XxxVo`
- Entities: `Xxx` (table name like `T_STAFF` becomes `TStaff`)
- Mappers: `XxxMapper`
- Methods: `camelCase` (e.g., `selectUserList`, `queryPageList`)
- Constants: `UPPER_SNAKE_CASE`

### Controller Patterns
```java
@Validated
@Slf4j
@RequiredArgsConstructor
@RestController
@RequestMapping("/system/user")
public class SysUserController extends BaseController {

    private final ISysUserService userService;

    // List query
    @GetMapping("/list")
    public TableDataInfo<SysUser> list(SysUser user, PageQuery pageQuery) {
        return userService.selectPageUserList(user, pageQuery);
    }

    // Get by ID
    @GetMapping("/{id}")
    public R<XxxVo> getInfo(@NotNull @PathVariable Long id) {
        return R.ok(service.queryById(id));
    }

    // Add
    @Log(title = "xxx", businessType = BusinessType.INSERT)
    @PostMapping()
    public R<Void> add(@Validated(AddGroup.class) @RequestBody XxxBo bo) {
        return toAjax(service.insertByBo(bo));
    }

    // Edit
    @Log(title = "xxx", businessType = BusinessType.UPDATE)
    @PutMapping()
    public R<Void> edit(@Validated(EditGroup.class) @RequestBody XxxBo bo) {
        return toAjax(service.updateByBo(bo));
    }

    // Delete
    @Log(title = "xxx", businessType = BusinessType.DELETE)
    @DeleteMapping("/{ids}")
    public R<Void> remove(@NotEmpty @PathVariable Long[] ids) {
        return toAjax(service.deleteWithValidByIds(Arrays.asList(ids), true));
    }
}
```

### Error Handling
- Use `R<T>` for unified response wrapper
- Use `AjaxResult` for legacy compatibility
- Use `@Log` annotation for operation logging
- Use `BusinessType` enum for log categorization
- Use `toAjax()` helper for CRUD operations
- Validation returns `R.fail()` with message

### Database
- Database: Dameng (DM) with MySQL compatibility mode
- ORM: MyBatis Plus
- Primary key strategy: `ASSIGN_ID` (snowflake)
- Logical delete: deleted=2, not_deleted=0
- Field strategy: `NOT_NULL` for insert/update
- Mapper XML location: `classpath*:mapper/**/*Mapper.xml`
- Entity scan: `com.jingwei.**.domain`

### Testing
- Framework: JUnit 5 (Jupiter)
- Use `@SpringBootTest` for integration tests
- Use `@DisplayName` for test description
- Use `@Disabled` to skip tests
- Use `@Timeout` for timeout
- Use `@RepeatedTest` for repeated execution
- Profile-based test execution: `-Plocal` `-Pdev` `-Pprod`
- Exclude tests tagged with `exclude`

### Logging
- Framework: SLF4J with Logback
- Configuration: `logback-plus.xml`
- Default level: `debug` for `com.jingwei`, `warn` for Spring
- Use `@Slf4j` annotation

### Security & Auth
- Framework: Sa-Token
- Token name: `Authorization`
- Token timeout: 86400 seconds (1 day)
- Activity timeout: 1800 seconds (30 minutes)
- Use `@SaCheckPermission("resource:action")` for permission control
- Use `@SaCheckLogin` for login verification
- Password encryption: BCrypt

### API Documentation
- Framework: SpringDoc (OpenAPI 3)
- Path: `/swagger-ui.html`
- Authorization: Bearer token in header
- Add `@Operation(summary = "...")` for endpoint documentation

## Database Migration Notes

Current: Dameng (DM) database
Target: PostgreSQL 17

Key differences to handle:
- Dameng uses VARCHAR(32) for IDs; PostgreSQL can use VARCHAR or UUID
- Dameng TIMESTAMP(0) vs PostgreSQL TIMESTAMP
- Dameng CHAR vs PostgreSQL CHAR
- Dameng specific functions need PostgreSQL equivalents
- Dameng cluster primary key syntax differs

SQL files location: `../script/sql/dm/`

## Configuration Files

- Main: `application.yml`
- Profiles: `application-dev.yml`, `application-prod.yml`
- Server port: 8089
- Context path: `/api`
