# TODO

## Reference data commands — unimplemented

Candidate `tess ref <resource> list` commands from `/ReferenceData/`, grouped and ranked by value for a performing-arts CRM workflow.

### High — supports likely next features

These support fundraising, membership, and ticketing workflows that are natural extensions of the existing CRM commands.

#### Fundraising / Development (19)

- [ ] AppealCategories — categories for fundraising appeals
- [ ] CampaignCategories — categories for fundraising campaigns
- [ ] ContributionDesignations — contribution designation codes
- [ ] ContributionImportSets — import set definitions for contributions
- [ ] DesignationCodes — codes for directing contributions
- [ ] DonationLevels — giving levels/tiers
- [ ] PhilanthropyTypes — types of philanthropic activity
- [ ] PlanPriorities — priority levels for development plans
- [ ] PlanSources — sources for development plans
- [ ] PlanStatuses — status values for development plans
- [ ] PlanTypes — types of development/solicitation plans
- [ ] PlannedGivingCodes — codes for planned giving
- [ ] PlannedGivingFundings — funding types for planned gifts
- [ ] PlannedGivingGiftTypes — types of planned gifts
- [ ] PlannedGivingOnFiles — on-file statuses for planned gifts
- [ ] PlannedGivingPurposes — purposes for planned gifts
- [ ] PlannedGivingSources — sources for planned gifts
- [ ] PlannedGivingStatuses — statuses for planned gifts
- [ ] RecognitionTypes — donor recognition type definitions

#### Membership (7)

- [ ] MembershipBenefitFrequencies — frequency options for membership benefits
- [ ] MembershipBenefitTypes — types of membership benefits
- [ ] MembershipLevelCategories — categories for membership levels
- [ ] MembershipLevelTrends — trend classifications for membership levels
- [ ] MembershipPeriods — membership period definitions
- [ ] MembershipStandings — standing types for memberships
- [ ] MembershipStatuses — status values for memberships

#### Constituent / CRM (14)

Extends the existing constituent model with contact-point, relationship, and classification lookups.

- [ ] ConstituentProtectionTypes — protection type classifications
- [ ] ConstituentTypeAffiliates — affiliate definitions for constituent types
- [ ] ContactLogActivityTypes — activity types for contact logs
- [ ] ContactPointCategories — categories for contact points
- [ ] ContactPointCategoryPurposes — contact point category-to-purpose mappings
- [ ] ContactPointPurposeCategories — purpose-to-category mappings
- [ ] ContactPointPurposes — purposes for contact points
- [ ] ContactTypes — types of contacts
- [ ] KeywordConstituentTypes — keyword-to-constituent-type mappings
- [ ] MailIndicators — mail indicator preferences
- [ ] NameStatuses — status values for constituent names
- [ ] PhoneIndicators — phone indicator preferences
- [ ] PhoneTypes — types of phone numbers
- [ ] RelationshipCategories — categories for constituent relationships

### Medium — supports ticketing, pricing, and marketing workflows

#### Ticketing / Pricing (14)

- [ ] DiscountTypes — types of ticket discounts
- [ ] EventLevels — event level classifications
- [ ] FeeCategories — categories for fees and surcharges
- [ ] HoldCodeCategories — categories for hold codes on seats/inventory
- [ ] PackageTypes — types of ticket packages
- [ ] PerformanceSegmentTypes — segment type classifications for performances
- [ ] PriceLayerTypes — price layer type definitions
- [ ] PriceTypeReasons — reason codes for price types
- [ ] PricingRuleCategories — categories for pricing rules
- [ ] PricingRuleMessageTypes — message types shown by pricing rules
- [ ] PricingRuleTypes — types of pricing rules
- [ ] SeasonTypes — types/classifications for seasons
- [ ] SubLineItemStatuses — status values for order sub-line items
- [ ] UpgradeCategories — categories for seat/ticket upgrades

#### Marketing / E-commerce (4)

- [ ] EmarketIndicators — e-marketing indicator definitions
- [ ] ResponseTypes — response type definitions for campaigns
- [ ] SourceGroups — source group classifications
- [ ] SourceTypes — source type definitions

#### Activities / Events (3)

- [ ] ActivityCategories — categories for activities
- [ ] AttendanceCategories — categories for attendance tracking
- [ ] InvitationStatuses — status values for invitations

#### Salutation / Naming (2)

- [ ] SalutationFormats — format definitions for salutations
- [ ] SalutationTypes — types of salutations

#### Shipping / Delivery (1)

- [ ] ShipMethods — shipping method definitions

### Low — niche or admin-only

#### Financial / Billing (8)

- [ ] AccountTypes — account type definitions
- [ ] BatchTypeGroups — groups for batch types
- [ ] BatchTypes — types of financial batches
- [ ] BillingSchedules — billing schedule definitions
- [ ] BillingTypes — types of billing
- [ ] CurrencyTypes — currency type definitions
- [ ] DirectDebitAccountTypes — account types for direct debit
- [ ] GLAccounts — general ledger account definitions

#### Payment Processing (4)

- [ ] CardReaderTypes — types of card readers
- [ ] PaymentGatewayTransactionTypes — transaction types for payment gateways
- [ ] PaymentMethodGroups — groups for payment methods
- [ ] TriPOSCloudConfigurations — TriPOS cloud configuration settings

#### Portfolio / Prospect (3)

- [ ] PortfolioCustomElements — custom element definitions for portfolios
- [ ] RankTypes — rank type definitions for prospects
- [ ] ResearchTypes — research type definitions for prospect research

#### Benefits / Qualifications (4)

- [ ] BenefitUsageContexts — contexts for benefit usage
- [ ] QualificationCategories — categories for qualifications
- [ ] Qualifications — qualification definitions
- [ ] SpecialRequestCategories — categories for special requests

#### Production / Artistic (5)

- [ ] Colors — color definitions
- [ ] Composers — composer reference data
- [ ] Eras — era/period classifications
- [ ] Premieres — premiere type definitions
- [ ] Voices — voice type definitions (for casting)

#### Facility / Scheduling (5)

- [ ] BookingCategories — categories for facility bookings
- [ ] ResourceCategories — categories for resources
- [ ] SchedulePatternTypes — pattern types for scheduling
- [ ] ScheduleTypes — types of schedules
- [ ] TimeSlots — time slot definitions

#### Workers / Staffing (4)

- [ ] CrediteeTypes — types of creditees for productions
- [ ] JobTypes — job type definitions
- [ ] WorkerRoles — role definitions for workers
- [ ] WorkerTypes — worker type classifications

#### Access Control / Scanning (7)

- [ ] AccessControlAreas — access control areas for venue entry
- [ ] AccessControlCustomizationPoints — customization points for access control
- [ ] AccessControlEntrances — entrance definitions for access control
- [ ] AccessControlProfiles — access control profile configurations
- [ ] BarcodeTypes — barcode type definitions for ticket scanning
- [ ] ExternalBarcodeSources — external barcode source definitions
- [ ] ScanTypes — types of scans (entry, exit, etc.)

#### Gift Aid (UK tax) (6)

Not applicable to the Folger.

- [ ] GiftAidContactMethods
- [ ] GiftAidDocumentStatuses
- [ ] GiftAidIneligibleReasons
- [ ] GiftAidRates
- [ ] GiftAidStatuses
- [ ] GiftAidTypes

#### Tax Receipting (2)

- [ ] CharitableTaxReceiptOrganizations — organizations for charitable tax receipts
- [ ] CharitableTaxReceiptVoidReasons — void reason codes for tax receipts

### Very low — system internals

Installation-specific or tied to desktop/web UI configuration.

#### Printing / Formats (3)

- [ ] Fonts
- [ ] Formats
- [ ] Printers

#### Templates / Design (5)

- [ ] DesignElementGroups
- [ ] DesignElements
- [ ] Designs
- [ ] TemplateCategories
- [ ] TemplateTypes

#### Web / TNEW / Wallet (7)

- [ ] TNEWCustomizationPoints
- [ ] TNEWCustomizations
- [ ] TNEWDynamicEmailContents
- [ ] WalletTemplateTypes
- [ ] WalletTemplates
- [ ] WebContentTypes
- [ ] SalesLayoutButtonTypes

#### Help System (6)

- [ ] CustomHelpDocuments
- [ ] CustomHelpLinks
- [ ] HelpDocumentTypes
- [ ] HelpDocuments
- [ ] HelpKeywords
- [ ] HelpLinks

#### System / Administrative (23)

- [ ] AlertEvents
- [ ] AnalyticsCubes
- [ ] AppScreenTexts
- [ ] AssetTypes
- [ ] AuditActions
- [ ] CalendarText
- [ ] ControlGroupUserGroups
- [ ] ControlGroups
- [ ] CustomDefaultCategories
- [ ] CustomDefaults
- [ ] CustomProcedures
- [ ] CustomScreenTypes
- [ ] CustomScreens
- [ ] DocumentCategories
- [ ] EmailProfiles
- [ ] ListCategories
- [ ] MonitorEvents

#### Miscellaneous (23)

- [ ] CriterionOperators
- [ ] IntegrationDefaults
- [ ] Integrations
- [ ] MediaTypes
- [ ] MenuItems
- [ ] ObjectPermissions
- [ ] Organizations
- [ ] Origins
- [ ] PostalCodeLookups
- [ ] ReceiptSettings
- [ ] ReferenceColumns
- [ ] ReferenceTableCategories
- [ ] ReferenceTableUserGroups
- [ ] ReferenceTables
- [ ] ReportUserGroups
- [ ] Schools
- [ ] SegmentationTypes
- [ ] ServiceResourceUserGroups
- [ ] ServiceResources
- [ ] StepPriorities
- [ ] StepTypes
- [ ] SurveyQuestions
- [ ] SystemDefaults
- [ ] ZoneGroups
