# Task Owner
---
## Dogan Akyurek
## akyurek.dogan.dgn@gmail.com
---
# Home24 | Website Analyzer

## 📦 Requirements

Gets the following inputs from the customer
- url
    - check validity of the url if not fail fast.
    - check the url accesibility if not fail fast.

Returns the following details
- response
    - What HTML version has the document?
    - What is the page title?
    - How many headings of what level are in the document?
    - How many internal and external links are in the document? Are there any inaccessible links and how many?
    - Does the page contain a login form?

---

## 📦 Project Layout

### ⏱ Milestones (avg hours)

| Task                                            | Duration |
|-------------------------------------------------|----------|
| Requirement analysis & architecture design      | 1h       |
| Core architecture implementation & edge cases   | 1h       |
| Service development                             | 2h       |
| Testing & documentation                         | 1h       |
| Deployment & hand-off                           | 0.5h     |

---

### ✅ Completed Tasks

- [x] Requirement gathering & edge case analysis
- [x] Project layout & git repository initialization
- [x] System design and architecture skeleton
- [x] MVP & dry run.
- [x] Implementation improvements
- [x] UI Implementation
- [ ] Unit test coverage check & increase
- [x] Project documentation
- [x] Project delivery to Home24 team

---

### ⚠️ Limitations

- Logging is implemented using uber/zap. Due to time constraint, I initialized it with the production configuration using the Sugared Logger. Most logs could be converted to DEBUG level for more detailed output.
- Different functionalities are implemented as plugins, which can be activated on demand via configuration (see main.go, lines 54–57), however this is just for demonstraion since plugins are unfortunately tighlty couple to a few structs due to time constraint.
- Asynchronous processing is demonstrated in a basic way to show concurrency is possible. With more time, the system could be made more efficient (see html_analyzer.go, lines 25–40). All analyzing process can be async and concurrent.
- Test cases were primarily performed on https://bbc.com since it is a feature-rich site. Larger websites were not pushed to avoid excessive processing time.
- Containerization is skipped.
- CI/CD pipeline is skipped.
- Metrics & improvement performance is skipped.

### ⚠️ Missing parts / Improvements
- Advanced Logging & Metrics: Implement detailed logging and metrics to measure processing times, identify bottlenecks, and monitor system health effectively.
- Concurrency & Architecture: Current concurrency can be improved. Consider an event-driven architecture to handle tasks more efficiently, monitor processing, and reduce false positives.
- Testing: Unit, end-to-end (E2E), and performance tests were not included due to time constraints. Adding these is essential for reliability and maintainability.
- Plugin Configuration: Make all plugins configurable through a configuration file or flagging system to allow selective execution and reduce unnecessary processing.
- Component Decoupling: Separate requesting, analyzing, and classification logic into independent, event-driven components to allow horizontal scaling and better resource management.
- Timeouts & Monitoring: Introduce more granular timeouts per task, collect website-specific metrics, and set up alerting systems to detect anomalies or failures.
- Frontend UX Enhancements: Improve the frontend to provide more informative feedback, guidance, and insights for users during analysis.
- Performance & Load Testing: Conduct performance and load testing to determine system limits and define clear service-level agreements (SLAs).
- Containerization & Orchestration: Improve Docker containerization and consider Kubernetes deployment for scalability, fault tolerance, and easy management.
- Security: Strengthen security measures, including input validation, TLS verification, rate limiting, and safe handling of external URLs.
- Feature Expansion: Extend the scope of analysis to cover SEO, accessibility, and other advanced webpage features based on business requirements rather than focusing on a limited set of elements.
- Instead of a service/app implementation this problem can be solved as data pipeline approach in a better manner I think.

---

## 🚀 Getting Started

### 🧰 Prerequisites

- Golang must be installed on local system
- You can check the task.pdf in the source code for task details.

---

## 🧪 Running the Project

### Step 1 – Setup

- Golang must be installed in the system.

---

### Step 2 – Run Options

- Please use the make file command to run the project `make run`
- Following steps will be done automatically when you use the make file however better to check
    - Please copy the `.env.dist file as .env` on your local.
    - Please copy the `.config.yml.dist as config.yml` on your local
- The web page must be opened automatically otherwise please visit the `localhost:8080` if it's not due to your browser or local security settings please visit manually.
---

### Step 3 – Test the Application

- You can use following websites to check the application (simpler to extensive)
    - example.com -> Very clean so better to check simple steps.
    - bbc.com -> I used for local tests, pretty packed website.

## 📸 Screenrecordings

https://github.com/user-attachments/assets/f1c1b7b0-441c-4684-944e-59da70f6fdfe

https://github.com/user-attachments/assets/55685665-46fb-4ed2-81da-026b2cd7987e

https://github.com/user-attachments/assets/0aea6257-9891-4dd1-985f-76c6b7e44c82

## 🤝 Contributing

This is a sample project submitted to Home24.

---

## 📬 Contact
---

*Dogan Akyurek*
*akyurek.dogan.dgn@gmail.com*

For questions or feedback, feel free to open an issue or reach out via *akyurek.dogan.dgn@gmail.com*

---
